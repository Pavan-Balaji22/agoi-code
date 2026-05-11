package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"ai-agent/llm"
)

var llmClient *llm.OllamaClient = llm.NewOllamaClient("deepseek-r1:7b")

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func processMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	message := r.FormValue("message")
	if message == "" {
		http.Error(w, "Message is required", http.StatusBadRequest)
		return
	}
	llmClient.CallLLM(message)

	fmt.Fprintf(w, "Message processed: %s", message)
}

func Serve() {
	apiv1 := http.NewServeMux()
	apiv1.HandleFunc("/greet", greet)
	apiv1.HandleFunc("/process", processMessage)
	log.Println("Starting API server on :8080")
	log.Fatal(
		http.ListenAndServe(":8080", apiv1))
}
