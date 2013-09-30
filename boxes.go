package textgui

/*
 * Functions and types for drawing ASCII frames and dealing with boxes
 * and rectangles.
 *
 */

const (
	TLCHAR = "┌" // top left
	TRCHAR = "┐" // top right
	BLCHAR = "└" // bottom left
	BRCHAR = "┘" // bottom right
	VCHAR  = "│" // vertical line, left side
	VCHAR2 = "│" // vertical line, right side
	HCHAR  = "─" // horizontal line
	HCHAR2 = "─" // horizontal bottom line
)

type Rect struct {
	x int
	y int
	w int
	h int
}

type Box struct {
	frame *Rect // The rectangle around the box, for placement
	inner *Rect // The rectangle inside the box, for content
}

func NewBox() *Box {
	return &Box{&Rect{0, 0, 0, 0}, &Rect{0, 0, 0, 0}}
}

func (r *Box) center(container *Box) {
	widthleftover := container.inner.w - r.frame.w
	heightleftover := container.inner.h - r.frame.h
	r.frame.x = container.inner.x + widthleftover/2
	r.frame.y = container.inner.y + heightleftover/2
}

func (r *Box) fill(container *Box) {
	r.frame.x = container.inner.x
	r.frame.y = container.inner.y
	r.frame.w = container.inner.w
	r.frame.h = container.inner.h
}

func (r *Box) fillWithMargins(container *Box, margins int) {
	r.fill(container)
	r.frame.x += margins
	r.frame.y += margins
	r.frame.w -= margins * 2
	r.frame.h -= margins * 2
}

func (r *Box) fillWithPercentageMargins(container *Box, horizmarginp float32, vertmarginp float32) {
	horizmargin := int(float32(container.inner.w) * horizmarginp)
	vertmargin := int(float32(container.inner.h) * vertmarginp)
	r.fill(container)
	r.frame.x += horizmargin
	r.frame.y += vertmargin
	r.frame.w -= horizmargin * 2
	r.frame.h -= vertmargin * 2
}

func (r *Box) getContentPos() (int, int) {
	return r.inner.x, r.inner.y
}

func (r *Box) thirdSize(container *Box) {
	r.frame.w = container.inner.w / 3
	r.frame.h = container.inner.h / 3
}

func (r *Box) thirdPlace(container *Box) {
	r.frame.x = container.inner.x + container.inner.w/3
	r.frame.y = container.inner.y + container.inner.h/3
}

func (r *Box) nicePlacement(container *Box) {
	r.frame.x = container.inner.x
	r.frame.y = container.inner.y
	leftoverwidth := container.inner.w - r.frame.w
	leftoverheight := container.inner.h - r.frame.h
	if leftoverwidth > r.frame.w {
		r.frame.x += leftoverwidth / 3
	}
	if leftoverheight > r.frame.h {
		r.frame.y += leftoverheight / 3
	}
}

func (b *Box) place(container *Box) {
	b.frame.x = container.inner.x
	b.frame.y = container.inner.y
}

func (b *Box) GetInner() *Rect {
	return b.inner
}

func (b *Box) GetFrame() *Rect {
	return b.frame
}

func (b *Box) SetInner(r *Rect) {
	b.inner = r
}

func (b *Box) SetFrame(r *Rect) {
	b.frame = r
}
