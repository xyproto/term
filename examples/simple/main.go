package main

import (
	. "github.com/xyproto/textgui"
)

func main() {
	Init()
	Clear()
	Say(10, 7, "hi")
	Flush()
	WaitForKey()
	Close()
}
