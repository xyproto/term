Term
====

[![Build Status](https://travis-ci.org/xyproto/term.svg?branch=master)](https://travis-ci.org/xyproto/term)

Archived
--------

Please consider using these Go packages instead, which should cover the same functionality:

* [vt100](https://github.com/xyproto/vt100) for coloring text, reading keys or drawing terminal "graphics".
* [textoutput](https://github.com/xyproto/textoutput) for coloring text, with tags like `<blue>`.
* [readline](https://github.com/chzyer/readline) for a pure-Go implementation of similar functionality to the C `readline` library.

Online API documentation
------------------------

[API documentation at godoc](http://godoc.org/github.com/xyproto/term)

Features and limitations
------------------------

* Provides an easy way to get started drawing colorful characters at any position (X,Y) in a terminal.
* Uses ncurses and the [gdamore/tcell](https://github.com/gdamore/tcell) module.
* Is simple to use, but [vt100](https://github.com/xyproto/vt100) may be more complete.

Simple example
--------------

~~~go
package main

import (
	"github.com/xyproto/term"
)

func main() {
	term.Init()
	term.Clear()
	term.Say(10, 7, "hi")
	term.Flush()
	term.WaitForKey()
	term.Close()
}
~~~

General information
-------------------

* License: MIT
* Author: Alexander F. Rødseth &lt;rodseth@gmail.com&gt;
* Version: 0.3.0
