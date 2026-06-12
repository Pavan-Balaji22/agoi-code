package main

import (
	co "agoi-code/core"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello! Welcome to agoi code")
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("Please chose a model ")
	fmt.Printf(":> ")
	model, err := reader.ReadString('\n')
	llmClient := co.NewOllamaClient(model)
	checkError(err)
	response := make(chan string)
	for {
		fmt.Printf(":> ")
		input, err := reader.ReadString('\n')
		checkError(err)
		// res, err := llmClient.CallLLM(input)
		go llmClient.AsyncCallLLM(input, response)
		checkError(err)
		fmt.Printf("\n%s : %s \n", strings.TrimSpace(model), <-response)

	}

}
