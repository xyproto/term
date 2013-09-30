package main

import (
	"github.com/nsf/termbox-go"
	"github.com/xyproto/textgui"
)

func main() {
	textgui.Init()
	textgui.Clear()
	textgui.Write(10, 10, "hi", termbox.ColorWhite, termbox.ColorGreen|termbox.AttrBold)
	textgui.SetFg(termbox.ColorRed)
	textgui.DrawRaw(20, 12, ",,..----..,,\n|          |\n\\__________/")
	textgui.Flush()
	textgui.WaitForKey()
	textgui.Close()
}
