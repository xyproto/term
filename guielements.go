package term

func DrawBackground() {
	Clear()
}

func DrawBox(r *Box, extrude bool) *Rect {
	x := r.frame.X
	y := r.frame.Y
	width := r.frame.W
	height := r.frame.H
	FG1 := BOXLIGHT
	FG2 := BOXDARK
	if !extrude {
		FG1 = BOXDARK
		FG2 = BOXLIGHT
	}
	Write(x, y, TLCHAR, FG1, BOXBG)
	Write(x+1, y, Repeat(HCHAR, width-2), FG1, BOXBG)
	Write(x+width-1, y, TRCHAR, FG1, BOXBG)
	for i := y + 1; i < y+height; i++ {
		Write(x, i, VCHAR, FG1, BOXBG)
		Write(x+1, i, Repeat(" ", width-2), FG1, BOXBG)
		Write(x+width-1, i, VCHAR2, FG2, BOXBG)
	}
	Write(x, y+height-1, BLCHAR, FG1, BOXBG)
	Write(x+1, y+height-1, Repeat(HCHAR2, width-2), FG2, BOXBG)
	Write(x+width-1, y+height-1, BRCHAR, FG2, BOXBG)
	return &Rect{x, y, width, height}
}

func DrawList(r *Box, items []string, selected int) {
	for i, s := range items {
		color := LISTTEXT
		if i == selected {
			color = LISTFOCUS
		}
		Write(r.frame.X, r.frame.Y+i, s, color, BOXBG)
	}
}

func DrawButton(x, y int, text string, active bool) {
	color := BUTTONTEXT
	if active {
		color = BUTTONFOCUS
	}
	Write(x, y, "<  ", color, BG)
	Write(x+3, y, text, color, BG)
	Write(x+3+len(text), y, "  >", color, BG)
}

func DrawAsciiArt(x, y int, text string) int {
	var i int
	for i, line := range Splitlines(text) {
		Write(x, y+i, line, TEXTCOLOR, BOXBG)
	}
	return y + i
}

func DrawRaw(x, y int, text string) int {
	var i int
	for i, line := range Splitlines(text) {
		Write(x, y+i, line, TEXTCOLOR, BG)
	}
	return y + i
}
