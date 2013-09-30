package textgui

/*
 * Various functions for using curses/termbox
 *
 */

import (
	"github.com/nsf/termbox-go"
)

var (
	BG          = termbox.ColorBlue
	TITLECOLOR  = termbox.ColorCyan | termbox.AttrBold
	TEXTCOLOR   = termbox.ColorBlack
	BOXBG       = termbox.ColorWhite
	BOXLIGHT    = termbox.ColorWhite | termbox.AttrBold
	BOXDARK     = termbox.ColorBlack
	BUTTONFOCUS = termbox.ColorYellow | termbox.AttrBold
	BUTTONTEXT  = termbox.ColorWhite | termbox.AttrBold
	LISTFOCUS   = termbox.ColorRed
	LISTTEXT    = termbox.ColorBlack
)

func Write(x int, y int, text string, fg termbox.Attribute, bg termbox.Attribute) {
	pos := 0
	for _, r := range text {
		termbox.SetCell(x+pos, y, r, fg, bg)
		pos++
	}
}

func Say(x int, y int, text string) {
	pos := 0
	for _, r := range text {
		termbox.SetCell(x+pos, y, r, TEXTCOLOR, BG)
		pos++
	}
}

func Clear() {
	termbox.Clear(TEXTCOLOR, BG)
}

func ScreenWidth() int {
	return First(termbox.Size)
}

func ScreenHeight() int {
	return Second(termbox.Size)
}

func Flush() {
	termbox.Flush()
}

func Init() error {
	return termbox.Init()
}

func Close() {
	termbox.Close()
}

func PollEvent() *termbox.Event {
	e := termbox.PollEvent()
	return &e
}

// Wait for Esc or Enter to be pressed
func WaitForKey() {
	for {
		e := PollEvent()
		switch e.Type {
		case termbox.EventKey:
			switch e.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyEnter:
				return
			}
		}
	}
}

func SetFg(fg termbox.Attribute) {
	TEXTCOLOR = fg
}

func SetBg(bg termbox.Attribute) {
	BG = bg
}
