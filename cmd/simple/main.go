package main

import (
	"github.com/xyproto/term"
)

func main() {
	term.Init()
	term.Clear()
	term.Say(10, 7, "hi")
	term.Flush()
	term.WaitForKey()
	term.Close()
}
