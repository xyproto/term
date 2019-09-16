package main

import (
	"fmt"
	"github.com/xyproto/term"
	"strings"
)

// Loop and echo the input until "quit" is typed
func main() {
	fmt.Println(`Type "quit" when done.`)
	for {
		// Retrieve user input, with a prompt. Use ReadLn() for no prompt.
		line := term.Ask("> ")

		// Check if the user wants to quit
		switch line {
		case "quit", "end", "exit", "q":
			return
		}

		// Repeat what was just said
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			fmt.Println("I repeat: " + line)
		}
	}
}
