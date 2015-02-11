TextGUI
=======

[![Build Status](https://travis-ci.org/xyproto/textgui.svg?branch=master)](https://travis-ci.org/xyproto/textgui)

Text User Interface module for Go.
(and text can be graphics too?)

Online API documentation
------------------------

[textgui API documentation at godoc.org](http://godoc.org/github.com/xyproto/textgui)


Features and limitations
------------------------

* Uses ncurses and the termbox-go module.


Simple example
--------------

~~~ go
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
~~~

Another example
---------------

~~~ go
package main

import (
	"fmt"

	"github.com/xyproto/textgui"
)

// Loop and echo the input until "exit" is typed
func Taunt() {
	for {
		// Retrieve user input, with a prompt. Use ReadLn() for no prompt.
		line := textgui.Ask("> ")

		// Check if the user has had enough
		if line == "quit" {
			break
		}

		// Taunt endlessly
		fmt.Println("No, you are " + line + "!")
	}
}

func main() {
	fmt.Println(`
Welcome to Taunt 1.0!
	
Type "quit" when done.
	
Ready.

`)
	Taunt()
}
~~~

Author and licence
-------------------

* License: MIT
* Author: Alexander Rødseth

