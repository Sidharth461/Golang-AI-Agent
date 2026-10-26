// package main

// import (
// 	"bufio"
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"

// 	"github.com/joho/godotenv"
// 	"github.com/openai/openai-go/v2"
// 	"github.com/openai/openai-go/v2/option"
// )

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("err", err)
// 	}
// 	apiKey := os.Getenv("OPENROUTER_API_KEY")

// 	if apiKey == "" {
// 		log.Fatal("Api key is Required !")
// 	}
// 	ctx := context.Background()
// 	url := "https://openrouter.ai/api/v1"

// 	client := openai.NewClient(
// 		option.WithBaseURL(url),
// 		option.WithAPIKey(apiKey),
// 	)

// 	messages := []openai.ChatCompletionMessageParamUnion{}

// 	model := "apodex/apodex-1.1-mini:free"

// 	// messages = append(messages, openai.UserMessage("what model are you?"))

// 	params := openai.ChatCompletionNewParams{
// 		Model:    model,
// 		Messages: messages,
// 	}

// 	scanner := bufio.NewScanner(os.Stdin)

// 	for {

// 		fmt.Print("\n>")

// 		if !scanner.Scan() {
// 			break
// 		}
// 		input := strings.TrimSpace(scanner.Text())

// 		if input == "" {
// 			continue
// 		}

// 		if input == "clear" {
// 			fmt.Print("\033[H\033[2j")
// 			continue
// 		}

// 		params.Messages = append(params.Messages, openai.UserMessage(input))

// 		res, err := client.Chat.Completions.New(ctx, params)
// 		if err != nil {
// 			fmt.Printf("API Error: %v\n", err)
// 			continue
// 		}

// 		if len(res.Choices) == 0 {
// 			fmt.Println("Error: Received an empty response from OpenRouter. The model might be rate-limited or busy.")
// 			continue // Skips to the next terminal prompt safely
// 		}

// 		output := res.Choices[0].Message.Content
// 		fmt.Println(output)

// 		params.Messages = append(params.Messages, openai.AssistantMessage(output))

// 		count := len(params.Messages)

// 		fmt.Println(count)

// 		if count >= 10 {
// 			params.Messages = params.Messages[count-4:]
// 			fmt.Println("History Get Refreshed after the 10th Message!")
// 		}
// 	}

// 	fmt.Println("Golang AI Agent Is Running")

// }

package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/Sidharth461/handlers"
	"github.com/Sidharth461/internal/ai"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not loaded:", err)
	}

	aiClient, err := ai.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	chatHandler := &handlers.ChatHandler{
		AI: aiClient,
	}

	r := chi.NewRouter()

	// Chat API endpoint
	r.Post("/api/chat", chatHandler.ServeHTTP)
	r.Get("/api/chats", chatHandler.ServeHTTP)

	// Serve frontend files
	fileServer := http.FileServer(http.Dir("./web"))
	r.Handle("/*", fileServer)

	log.Println("Go AI Agent running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed:", err)
	}
}
