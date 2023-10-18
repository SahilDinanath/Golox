package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(len(os.Args))

	var numArgs = len(os.Args)
	switch numArgs {
	case 1:
	//TODO:
	//runPrompt()
	case 2:
		//TODO:
		//displayHelp()
	case 3:
		//TODO:
		//help()
	}
}
