TextGUI
=======

Text GUI module for Go. Uses ncurses and the termbox-go module.

* License: MIT
* Author: Alexander Rødseth <rodseth@gmail.com>


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

Online API documentation
------------------------

[go.pkgdoc.org](http://go.pkgdoc.org/github.com/xyproto/textgui)
