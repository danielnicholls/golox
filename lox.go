package main

import "os"
import "fmt"

func main() {
	if len(os.Args) {
		fmt.Println("Usage: golox [script]")
		os.Exit(64)
	} else if len(os.Args) == 1 {
		RunFile(os.Args[1])
	} else {
		RunPrompt()
	}
}
