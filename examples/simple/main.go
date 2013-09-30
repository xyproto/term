package main

import (
	"github.com/xyproto/textgui"
)

func main() {
	textgui.Init()
	textgui.Clear()
	textgui.Say(10, 7, "hi")
	textgui.Flush()
	textgui.WaitForKey()
	textgui.Close()
}
