# MEMOSGPT

This project integrates MEMOS, an open-source note-taking service, with OpenAI's powerful language model. It listens for webhooks triggered by new and updated notes, scans for questions prefixed with `/g`, queries the OpenAI API, and posts the responses as comments on the original note.

## Features

- **Webhook Listener**: Automatically triggers on new notes added in MEMOS.
- **Command Parsing**: Scans notes for the specific command `/g`.
- **OpenAI Integration**: Utilizes OpenAI's language model to generate answers for parsed questions.
- **Response Posting**: Automatically posts the generated answers back to the original note as a comment.

## Prerequisites

- Go (version 1.22.4 or newer)
- MEMOS instance with webhook capability
- OpenAI API key

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/morgansundqvist/memosgpt.git
   cd memosgpt
   ```

2. **Set up your environment variables:**
   Create a `.env` file in the project root and update it with your OpenAI API key and MEMOS webhook secret.

   ```plaintext
    OPENAI_API_KEY="your-openai-api-key-here"
    MEMOS_BASE_URL="http://localhost:5230"
    OPEN_AI_API_KEY="your-memos-api-key"
   ```

3. **Build the project:**

   ```bash
   go build
   ```

4. **Run the service:**
   ```bash
   ./memosgpt
   ```

## Configuration

- **Webhook Endpoint**: Ensure the webhook endpoint in MEMOS is set to the URL where this service is running, typically `http://your-ip:3214/wh`. Read more about webhooks in MEMOS here https://www.usememos.com/docs/advanced-settings/webhook

## Usage

Once the service is up and running, it will listen for incoming webhooks from MEMOS. When a note containing the `/g <question>` syntax is detected, it will:

1. Parse the question.
2. Send it to the OpenAI API.
3. Retrieve the response and post it back as a comment to the note.

Example Note:

```
Here is a question about physics:
/g What is the theory of relativity?
```

## Contributing

Contributions are welcome! If you'd like to improve the MEMOSGPT, please fork the repository and submit a pull request.

## License

Distributed under the MIT License.

## Acknowledgments

- [MEMOS](https://github.com/memos-org/memos) for the open-source note-taking platform.
- [go-openai](https://pkg.go.dev/github.com/sashabaranov/go-openai) for providing the Golang client for OpenAI.
- [grequests](https://github.com/levigross/grequests) for prividing a simple HTTP request package
