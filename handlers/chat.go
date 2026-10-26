package handlers

import (
	"encoding/json"
	"net/http"
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

func (chatAi *ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		chatAi.mu.RLock()
		records := make([]ChatRecord, len(chatAi.records))
		copy(records, chatAi.records)
		chatAi.mu.RUnlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(records)
		return

	case http.MethodPost:
		// Continue with your existing POST logic below.

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ChatRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Req Body", http.StatusBadRequest)
		return
	}

	if req.Message == "" {
		http.Error(w, "Message Is required", http.StatusBadRequest)
		return
	}

	messages := make([]openai.ChatCompletionMessageParamUnion, 0, len(req.History)+1)

	for _, msg := range req.History {
		switch msg.Role {
		case "user":
			messages = append(messages, openai.UserMessage(msg.Content))
		case "assistant":
			messages = append(messages, openai.AssistantMessage(msg.Content))
		}
	}

	messages = append(messages, openai.UserMessage(req.Message))

	resReply, err := chatAi.AI.Chat(r.Context(), messages)
	if err != nil {
		http.Error(w, "failed to get res from ai", http.StatusBadGateway)
		return
	}

	// Save the user message and AI response.
	now := time.Now()

	chatAi.mu.Lock()
	chatAi.records = append(chatAi.records,
		ChatRecord{
			Role:      "user",
			Content:   req.Message,
			CreatedAt: now,
		},
		ChatRecord{
			Role:      "assistant",
			Content:   resReply,
			CreatedAt: now,
		},
	)
	chatAi.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ChatResponse{
		Reply: resReply,
	})
}
