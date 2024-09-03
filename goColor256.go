package GoColor

import "strconv"

const (
	foreground256 = "\033[38;5;"
	background256 = "\033[48;5;"
)

func makeForeground256EscapeSequence(color int) string {
	return foreground256 + strconv.Itoa(color) + "m"
}

func makeBackground256EscapeSequence(color int) string {
	return background256 + strconv.Itoa(color) + "m"
}

// Export functions

// SetForegroundColor256 sets the foreground color of console output until the next Reset call
//
// Param color: foreground color to be set, ranging from 0 to 255 (extended ANSI palette)
func SetForegroundColor256(color int) {
	if color < 0 || color > 255 {
		panic("ANSI color out of range")
	}
	print(makeForeground256EscapeSequence(color))
}

// SetBackgroundColor256 sets the background color of console output until the next Reset call
//
// Param color: background color to be set, ranging from 0 to 255 (extended ANSI palette)
func SetBackgroundColor256(color int) {
	if color < 0 || color > 255 {
		panic("ANSI color out of range")
	}
	print(makeBackground256EscapeSequence(color))
}

// Print256 formats the given string with ANSI color codes and outputs it to the console
//
// Param str: string to be formatted
//
// Param colorFg: foreground color of the text, ranging from 0 to 255 (extended ANSI palette)
//
// Param colorBg: background color of the text, ranging from 0 to 255 (extended ANSI palette)
func Print256(str string, colorFg int, colorBg int) {
	if colorFg < 0 || colorFg > 255 || colorBg < 0 || colorBg > 255 {
		panic("ANSI color out of range")
	}
	print(makeForeground256EscapeSequence(colorFg) + makeBackground256EscapeSequence(colorBg) + str + reset)
}

// Println256 formats the given string with ANSI color codes and outputs it to the console, appending it with the "\n" character
//
// Param str: string to be formatted
//
// Param colorFg: foreground color of the text, ranging from 0 to 255 (extended ANSI palette)
//
// Param colorBg: background color of the text, ranging from 0 to 255 (extended ANSI palette)
func Println256(str string, colorFg int, colorBg int) {
	if colorFg < 0 || colorFg > 255 || colorBg < 0 || colorBg > 255 {
		panic("ANSI color out of range")
	}
	println(makeForeground256EscapeSequence(colorFg) + makeBackground256EscapeSequence(colorBg) + str + reset)
}

// PrintFg256 formats the given string with ANSI color codes and outputs it to the console
//
// Param str: string to be formatted
//
// Param color: foreground color of the text, ranging from 0 to 255 (extended ANSI palette)
func PrintFg256(str string, color int) {
	if color < 0 || color > 255 {
		panic("ANSI color out of range")
	}
	print(makeForeground256EscapeSequence(color) + str + reset)
}

// PrintlnFg256 formats the given string with ANSI color codes and outputs it to the console, appending it with the "\n" character
//
// Param str: string to be formatted
//
// Param color: foreground color of the text, ranging from 0 to 255 (extended ANSI palette)
func PrintlnFg256(str string, color int) {
	if color < 0 || color > 255 {
		panic("ANSI color out of range")
	}
	println(makeForeground256EscapeSequence(color) + str + reset)
}

// PrintBg256 formats the given string with ANSI color codes and outputs it to the console
//
// Param str: string to be formatted
//
// Param color: background color of the text, ranging from 0 to 255 (extended ANSI palette)
func PrintBg256(str string, color int) {
	if color < 0 || color > 255 {
		panic("ANSI color out of range")
	}
	print(makeBackground256EscapeSequence(color) + str + reset)
}

// PrintlnBg256 formats the given string with ANSI color codes and outputs it to the console, appending it with the "\n" character
//
// Param str: string to be formatted
//
// Param color: background color of the text, ranging from 0 to 255 (extended ANSI palette)
func PrintlnBg256(str string, color int) {
	if color < 0 || color > 255 {
		panic("ANSI color out of range")
	}
	println(makeBackground256EscapeSequence(color) + str + reset)
}
