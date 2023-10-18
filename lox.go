package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf(">")
	for scanner.Scan() {
		fmt.Printf(">")
		//TODO:
		//run(scanner.Text)
	}
}

func runFile(path string) {
	file, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(file))
	//TODO:
	//run(file)

}
func main() {
	var numArgs = len(os.Args)
	switch numArgs {
	case 1:
		//TODO:
		runPrompt()
	case 2:
		//TODO:
		runFile(os.Args[1])
	default:
		//TODO:
		//help()
	}
}
