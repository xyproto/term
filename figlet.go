package textgui

import (
	"github.com/getwe/figlet4go"
)

// Using the figlet4go library
func Figlet(msg string) (string, error) {
	ascii_render := figlet4go.NewAsciiRender()
	return ascii_render.Render(msg)
}
