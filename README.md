# Golang AI Agent

A full-stack AI chat application built with Go on the backend and React + Vite on the frontend. The app sends chat messages to an AI model through OpenRouter and displays the conversation in a modern chat interface.

## What this project includes

- Go backend server using `chi` for routing
- OpenRouter integration via the OpenAI-compatible Go client
- React frontend for chat UI and message sending
- Persistent in-memory chat records during runtime
- Support for conversation history in requests to the AI model
- API endpoint for chat interactions and frontend serving

## Tech stack

- Backend: Go, Chi router
- Frontend: React, Vite
- AI API: OpenRouter
- Environment management: `.env`

## Project structure

```text
.
├── main.go
├── go.mod
├── go.sum
├── .env.example
├── .env
├── handlers/
│   └── chat.go
├── internal/
│   └── ai/
│       └── client.go
├── frontend/
│   └── frontend/
│       ├── package.json
│       ├── src/
│       ├── public/
│       └── vite.config.js
└── README.md
```

## Features

- Send user prompts from the frontend
- Forward messages to the selected AI model
- Include prior chat history for context-aware responses
- Save chat records in the Go server memory
- New chat/reset support in the frontend
- Error handling for invalid requests and AI failures

## Environment setup

Create a `.env` file in the project root using the sample file as a guide:

```env
OPENROUTER_API_KEY=your_openrouter_api_key_here
OPENROUTER_URL=https://openrouter.ai/api/v1
MODEL_NAME=Your_Model_Name
```

You can copy the example values:

```bash
copy .env.example .env
```

Then update the values with your actual OpenRouter credentials.

## Running the app

### 1. Start the Go backend

```bash
go run .
```

This starts the server on:

```text
http://localhost:8080
```

### 2. Start the React frontend

```bash
cd frontend/frontend
npm install
npm run dev
```

Then open the Vite local URL in the browser, usually:

```text
http://localhost:5173
```

## API

### POST /api/chat

Request body:

```json
{
  "message": "Hello!",
  "history": [
    { "role": "user", "content": "Hi" },
    { "role": "assistant", "content": "Hello there!" }
  ]
}
```

Response:

```json
{
  "reply": "Hi! How can I help?"
}
```

### GET /api/chats

Returns the chat records stored in memory by the backend.

## Notes

- The current project stores chat history in memory only, so it resets when the server restarts.
- The AI model is configured through environment variables, making it easy to switch providers or models.
- This project is a great base for extending with authentication, persistent storage, streaming responses, and a richer multi-chat experience.

## License

This project is for personal learning and experimentation.
