package textgui

import (
	"github.com/getwe/figlet4go"
)

// Using the figlet4go library
func Figlet(msg string) (string, error) {
	// Figlet4go also support using figlet fonts (/usr/share/figlet/*.flf).
	// These also supports lowercase characters.
	// Look at slant.flf and big.flf, for instance.
	// TODO: Use big.flf, if available
	ascii_render := figlet4go.NewAsciiRender()
	return ascii_render.Render(msg)
}
