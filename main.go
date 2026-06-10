package main

import (
	"agoi-code/llm"
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
	llmClient := llm.NewOllamaClient(model)
	checkError(err)
	for {
		fmt.Printf(":> ")
		input, err := reader.ReadString('\n')
		checkError(err)
		res, err := llmClient.CallLLM(input)
		checkError(err)
		fmt.Printf("%s : %s \n", strings.Replace(model, "\n", "", 1), res)

	}

}
