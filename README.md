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
	. "github.com/xyproto/textgui"
)

func main() {
	Init()
	Clear()
	Say(10, 7, "hi")
	Flush()
	WaitForKey()
	Close()
}
```

Online API documentation
------------------------

[textgui API documentation at godoc.org](http://godoc.org/github.com/xyproto/textgui)
