package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"agoi-code/llm"
)

var llmClient *llm.OllamaClient = llm.NewOllamaClient("gemma4:31b-cloud")

type request struct {
	Message string `json:"message"`
}

func processMessage(w http.ResponseWriter, r *http.Request) {
	body := &request{}
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		http.Error(w, "Message is required", http.StatusBadRequest)
		return
	}
	llmClient.CallLLM(body.Message)

	fmt.Fprintf(w, "Message processed: %s", body.Message)
}

func Serve() {
	apiv1 := http.NewServeMux()
	apiv1.HandleFunc("POST /process", processMessage)
	main := http.NewServeMux()
	main.Handle("/apiv1/", http.StripPrefix("/apiv1", apiv1))
	log.Println("Starting API server on :8080")
	log.Fatal(
		http.ListenAndServe(":8080", main))
}
