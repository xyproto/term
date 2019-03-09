package main

import (
	"fmt"

	"github.com/xyproto/term"
)

// Loop and echo the input until "quit" is typed
func Repeat() {
OUT:
	for {
		// Retrieve user input, with a prompt. Use ReadLn() for no prompt.
		line := term.Ask("> ")

		// Check if the user wants to quit
		switch line {
		case "quit", "end", "exit", "q":
			break OUT
		}

		// Repeat what was just said
		fmt.Println("You said: " + line)
	}
}

func main() {
	fmt.Print(`
Welcome to Repeat 1.0!

Type "quit" when done.

Ready.

`)
	Repeat()
}
