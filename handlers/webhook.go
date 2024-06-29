package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/levigross/grequests"
	"github.com/morgansundqvist/memosgpt/config"
	"github.com/sashabaranov/go-openai"
)

type Memo struct {
	Name    string `json:"name"`
	Uid     string `json:"uid"`
	Content string `json:"content"`
}

type WebhookRequestBody struct {
	ActiityType string `json:"activityType"`
	Memo        Memo   `json:"memo"`
}

func HandleWebHook(c *fiber.Ctx) error {
	var body WebhookRequestBody

	if err := c.BodyParser(&body); err != nil {
		fmt.Printf("Error parsing webhook body: %v\n", err)
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if body.ActiityType != "memos.memo.updated" && body.ActiityType != "memos.memo.created" {
		fmt.Printf("Ignoring activity type: %s\n", body.ActiityType)
		return c.SendStatus(200)
	}

	go AskOpenAI(body.Memo)

	return c.SendStatus(200)
}

func AskOpenAI(memo Memo) {
	configInstance := config.GetConfig()
	// split body.Memo.content by \n
	contentRows := strings.Split(memo.Content, "\n")
	response := ""
	for _, row := range contentRows {
		if strings.HasPrefix(row, "/g") {
			println("Found a command")
			//remove /g from the row
			question := strings.TrimPrefix(row, "/g")
			question = strings.TrimLeft(question, " ")
			println("Question: ", question)
			client := openai.NewClient(configInstance.OpenAIAPIKey)
			resp, err := client.CreateChatCompletion(
				context.Background(),
				openai.ChatCompletionRequest{
					Model: openai.GPT4o,
					Messages: []openai.ChatCompletionMessage{
						{
							Role:    openai.ChatMessageRoleSystem,
							Content: "The following is a conversation with an AI assistant. The assistant is helpful, creative, clever, and very friendly. And you have an understanding of the following context " + memo.Content,
						},
						{
							Role:    openai.ChatMessageRoleUser,
							Content: question,
						},
					},
				},
			)
			if err != nil {
				fmt.Printf("ChatCompletion error: %v\n", err)
			}

			response = resp.Choices[0].Message.Content

			memoCommentCreateRequest := MemoCommentCreateRequest{
				Content: response,
			}

			memoResponse, err := grequests.Post(configInstance.MemosBaseURL+"/api/v1/"+memo.Name+"/comments", &grequests.RequestOptions{JSON: memoCommentCreateRequest, Headers: map[string]string{"Authorization": "Bearer " + configInstance.MemosAPIKey}})
			if err != nil {
				fmt.Printf("Error posting response to memo: %v\n", err)
			}

			if memoResponse.StatusCode != 200 {
				fmt.Printf("Error posting response to memo: %v\n", memoResponse.String())
			}

		}

	}
}

type MemoCommentCreateRequest struct {
	Content string `json:"content"`
}
