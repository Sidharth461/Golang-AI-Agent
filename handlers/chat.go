package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/Sidharth461/internal/ai"
	"github.com/openai/openai-go/v2"
)

type ChatHandler struct {
	AI *ai.Client

	mu      sync.RWMutex
	records []ChatRecord
}

type ChatRecord struct {
	Role      string    `json:"role"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type ChatRequest struct {
	Message string        `json:"message"`
	History []chatMessage `json:"history"`
}

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Reply string `json:"reply"`
}

func (chatAI *ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:

		chatAI.mu.RLock()
		records := make([]ChatRecord, len(chatAI.records))
		copy(records, chatAI.records)
		chatAI.mu.RUnlock()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(records); err != nil {
			log.Printf("Failed to encode chat records: %v", err)
		}
		return

	case http.MethodPost:
		// POST: Receive a message and get an AI reply.

		var req ChatRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// 2. Validate the message.
		req.Message = strings.TrimSpace(req.Message)

		if req.Message == "" {
			http.Error(w, "Message is required", http.StatusBadRequest)
			return
		}

		//  Convert the conversation history into AI messages.
		messages := make([]openai.ChatCompletionMessageParamUnion, 0, len(req.History)+1)

		for _, msg := range req.History {
			switch msg.Role {
			case "user":
				messages = append(messages, openai.UserMessage(msg.Content))
			case "assistant":
				messages = append(messages, openai.AssistantMessage(msg.Content))
			}
		}

		// Add the current user message.
		messages = append(messages, openai.UserMessage(req.Message))

		// Ask the AI to generate a reply.
		if chatAI.AI == nil {
			http.Error(w, "AI client is not configured", http.StatusInternalServerError)
			return
		}

		reply, err := chatAI.AI.Chat(r.Context(), messages)
		if err != nil {
			log.Printf("AI request failed: %v", err)
			http.Error(w, "Failed to get response from AI", http.StatusBadGateway)
			return
		}

		// Save both messages in memory.
		now := time.Now()

		chatAI.mu.Lock()
		chatAI.records = append(chatAI.records, ChatRecord{Role: "user", Content: req.Message, CreatedAt: now}, ChatRecord{Role: "assistant", Content: reply, CreatedAt: now})
		chatAI.mu.Unlock()

		//  Return the AI reply to the frontend.
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(ChatResponse{Reply: reply}); err != nil {
			log.Printf("Failed to encode AI response: %v", err)
		}
		return

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
