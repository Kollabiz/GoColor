# GoColor

A small library to output colored text to the console<br>
It uses ANSI color codes to output 24 bit (if available) deep colored text, and also supports <b>bold</b> and <u>underline</u> text<br>
<br>
# Usage

### Install

To install the library, just copy the following command to your command line<br>
```go get github.com/Kollabiz/GoColor```<br>
<br>

### Examples

The <b>GoColor</b> package has a set of functions that behave similarly to the built-in <b>fmt</b> library.<br>
The naming of the functions is similar, with <b>...print...()</b> functions outputting ANSI formatted strings and <b>...println...()</b> functions adding a "\n" character to the end of the output string.<br>
<br>
For example:<br>

```go
package main

import (
	"GoColor"
	"image/color"
)

func main() {
	// The PrintlnFg function writes a string formatted with ANSI background color
	// code to the console, appending it with "\n"
	GoColor.PrintlnFg("Hello, World!", color.RGBA{R:168, G:97, B:206})
}
```

Also, <b>GoColor</b> module has <b>Set...()</b> functions, which set color and bold/underline styles globally, until the next <b>Reset</b> function is called (or the color is overriden)<br>

```go
package main

import (
	"GoColor"
	"fmt"
	"image/color"
)

func main() {
	GoColor.SetBold()
	GoColor.SetForegroundColor(color.RGBA{R:168, G:97, B:206})
	// Now, everything outputted to the console will be bold and have the set foreground color
	fmt.Println("Hello, World!")
	// The color will remain until the next Reset() call, or until it is overriden with another GoColor function
	// Note: All Print...() functions add a reset sequence to the end of the string
	GoColor.PrintBg("Uh oh, the color got overriden!", color.RGBA{R:78,G:211,B:191})
	// Also, you can reset the bold/underline separately from color using respective functions
	GoColor.ResetBold()
}
```
