textgui
=======

Text GUI module for Go, uses ncurses

MIT license
Alexander Rødseth 2013

Simple example
--------------

```go
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
```
