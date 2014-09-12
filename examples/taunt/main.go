package main

import (
	"fmt"

	"github.com/xyproto/textgui"
)

// Loop and echo the input until "exit" is typed
func Taunt() {
	for {
		// Retrieve user input, with a prompt. Use ReadLn() for no prompt.
		line := textgui.Ask("> ")

		// Check if the user has had enough
		if line == "quit" {
			break
		}

		// Taunt endlessly
		fmt.Println("No, you are " + line + "!")
	}
}

func main() {
	fmt.Println(`
Welcome to Taunt 1.0!
	
Type "quit" when done.
	
Ready.

`)
	Taunt()
}
