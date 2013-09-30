package textgui

import (
	"bufio"
	"fmt"
	"os"
)

/*
 * Various helper functions
 *
 */

type ReturnsTwoInts func() (int, int)

// Get the first value from the function that returns two ints
func First(f ReturnsTwoInts) int {
	v, _ := f()
	return v
}

// Get the second value from the function that returns two ints
func Second(f ReturnsTwoInts) int {
	_, v := f()
	return v
}

// Read a line from stdin
func ReadLn() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return line[:len(line)-1]
}

// Ask a question, wait for textual input followed by a newline
func Ask(prompt string) string {
	fmt.Print(prompt)
	return ReadLn()
}

// Ask a yes/no question, don't wait for newline
func AskYesNo(question string, noIsDefault bool) bool {
	var s string
	alternatives := "Yn"
	if noIsDefault {
		alternatives = "yN"
	}
	fmt.Printf(question + " [" + alternatives + "] ")
	fmt.Scanf("%s", &s)
	if noIsDefault {
		// Anything that isn't yes is "no" (false)
		return s == "Y" || s == "y"
	}
	// Anything that isn't no is "yes" (true)
	return !(s == "N" || s == "n")
}
