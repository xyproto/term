package main

import (
	. "github.com/xyproto/term"
)

func main() {
	Init()
	Clear()

	screen := NewBox()

	frame := screen.GetFrame()
	frame.W = ScreenWidth()
	frame.H = ScreenHeight()
	screen.SetFrame(frame)

	inner := screen.GetInner()
	inner.X = 0
	inner.Y = 3 // This space is used by the title
	inner.W = frame.W - inner.X
	inner.H = frame.H - inner.Y
	screen.SetInner(inner)

	infoBox := NewBox()
	infoBox.SetThirdSize(screen)
	infoBox.FillWithPercentageMargins(screen, 0.07, 0.1)
	infoBox.SetInner(DrawBox(infoBox, true))

	Flush()
	WaitForKey()
	Close()
}
