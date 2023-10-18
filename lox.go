package main

import (
	"bufio"
	"fmt"
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
func main() {
	var numArgs = len(os.Args)
	switch numArgs {
	case 1:
		//TODO:
		runPrompt()
	case 2:
		//TODO:
		//displayHelp()
	default:
		//TODO:
		//help()
	}
}
