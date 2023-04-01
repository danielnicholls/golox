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

func RunFile(path string) {
	b, e := os.ReadFile(path)
	if e != nil {
		panic(e)
	}
	Run(b)
}
