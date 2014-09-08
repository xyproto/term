TextGUI
=======

[![Build Status](https://travis-ci.org/xyproto/textgui.svg?branch=master)](https://travis-ci.org/xyproto/textgui)

Text User Interface module for Go. ("Text/Graphical User Interface")

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
~~~

Author and licence
-------------------

* License: MIT
* Author: Alexander Rødseth <rodseth@gmail.com>


