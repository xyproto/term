package textgui

/*
 * Functions and types for drawing ASCII frames and dealing wi.H boxes
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
	X int
	Y int
	W int
	H int
}

type Box struct {
	frame *Rect // The rectangle around the box, for placement
	inner *Rect // The rectangle inside the box, for content
}

func NewBox() *Box {
	return &Box{&Rect{0, 0, 0, 0}, &Rect{0, 0, 0, 0}}
}

func (r *Box) center(container *Box) {
	widthleftover := container.inner.W - r.frame.W
	heightleftover := container.inner.H - r.frame.H
	r.frame.X = container.inner.X + widthleftover/2
	r.frame.Y = container.inner.Y + heightleftover/2
}

func (r *Box) fill(container *Box) {
	r.frame.X = container.inner.X
	r.frame.Y = container.inner.Y
	r.frame.W = container.inner.W
	r.frame.H = container.inner.H
}

func (r *Box) fillWithMargins(container *Box, margins int) {
	r.fill(container)
	r.frame.X += margins
	r.frame.Y += margins
	r.frame.W -= margins * 2
	r.frame.H -= margins * 2
}

func (r *Box) fillWithPercentageMargins(container *Box, horizmarginp float32, vertmarginp float32) {
	horizmargin := int(float32(container.inner.W) * horizmarginp)
	vertmargin := int(float32(container.inner.H) * vertmarginp)
	r.fill(container)
	r.frame.X += horizmargin
	r.frame.Y += vertmargin
	r.frame.W -= horizmargin * 2
	r.frame.H -= vertmargin * 2
}

func (r *Box) getContentPos() (int, int) {
	return r.inner.X, r.inner.Y
}

func (r *Box) thirdSize(container *Box) {
	r.frame.W = container.inner.W / 3
	r.frame.H = container.inner.H / 3
}

func (r *Box) thirdPlace(container *Box) {
	r.frame.X = container.inner.X + container.inner.W/3
	r.frame.Y = container.inner.Y + container.inner.H/3
}

func (r *Box) nicePlacement(container *Box) {
	r.frame.X = container.inner.X
	r.frame.Y = container.inner.Y
	leftoverwidth := container.inner.W - r.frame.W
	leftoverheight := container.inner.H - r.frame.H
	if leftoverwidth > r.frame.W {
		r.frame.X += leftoverwidth / 3
	}
	if leftoverheight > r.frame.H {
		r.frame.Y += leftoverheight / 3
	}
}

func (b *Box) place(container *Box) {
	b.frame.X = container.inner.X
	b.frame.Y = container.inner.Y
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
