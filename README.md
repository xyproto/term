TextGUI
=======

Text GUI module for Go. Uses ncurses and the termbox-go module.

* License: MIT
* Author: Alexander Rødseth <rodseth@gmail.com>


A simple example
----------------

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

Another simple example
----------------------

```go
package main

import (
	"fmt"
	. "github.com/xyproto/textgui"
)

// Loop and echo the input until "exit" is typed
func taunt() {
	for {
		// The Ask function can be used instead of the two following lines
		fmt.Print("> ")
		line := ReadLn()

		// Check if the user wants to quit
		if line == "exit" {
			break
		}

		// Taunt endlessly
		fmt.Println("You are " + line + "!")
	}
}

func main() {
	fmt.Println()
	fmt.Println("Welcome to Taunt v.1.0")
	fmt.Println()
	fmt.Println("Type exit when done.")
	fmt.Println()
	fmt.Println("Ready.")
	taunt()
}
```

Online API documentation
------------------------

[textgui API documentation at godoc.org](http://godoc.org/github.com/xyproto/textgui)
