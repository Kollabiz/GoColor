package GoColor

import (
	"image/color"
	"strconv"
)

const (
	escSequence    = "\033["
	reset          = escSequence + "0m"
	bold           = escSequence + "1m"
	resetBold      = escSequence + "22m"
	underline      = escSequence + "4m"
	resetUnderline = escSequence + "24m"
	foreground     = "38"
	background     = "48"
)

// Colors

func makeForegroundColorEscapeSequence(color color.Color) string {
	r, g, b, _ := color.RGBA()
	return escSequence + foreground + ";2;" + strconv.Itoa(int(r/257)) + ";" + strconv.Itoa(int(g/257)) + ";" + strconv.Itoa(int(b/256)) + "m"
}

func makeBackgroundColorEscapeSequence(color color.Color) string {
	r, g, b, _ := color.RGBA()
	return escSequence + background + ";2;" + strconv.Itoa(int(r/257)) + ";" + strconv.Itoa(int(g/257)) + ";" + strconv.Itoa(int(b/256)) + "m"
}

// Export functions

// Print formats the given string with ANSI color codes and outputs it to the console
//
// Param str: string to be formatted
//
// Param colorFg: foreground color of the text
//
// Param colorBg: background color of the text
//
// Note: Even if your console does not support RGB color codes, it will still select the closest color from its palette
func Print(str string, colorFg color.Color, colorBg color.Color) {
	print(makeBackgroundColorEscapeSequence(colorBg) + makeForegroundColorEscapeSequence(colorFg) + str + reset)
}

// PrintFg formats the given string with ANSI foreground color code and outputs it to the console
//
// Param str: string to be formatted
//
// Param color: foreground color of the text
//
// Note: Even if your console does not support RGB color codes, it will still select the closest color from its palette
func PrintFg(str string, color color.Color) {
	print(makeForegroundColorEscapeSequence(color) + str + reset)
}

// PrintBg formats the given string with ANSI background color code and outputs it to the console
//
// Param str: string to be formatted
//
// Param color: background color of the text
//
// Note: Even if your console does not support RGB color codes, it will still select the closest color from its palette
func PrintBg(str string, color color.Color) {
	print(makeBackgroundColorEscapeSequence(color) + str + reset)
}

// PrintBold formats the given string with ANSI bold code and outputs it to the console
//
// Param str: string to be formatted
func PrintBold(str string) {
	print(bold + str + resetBold)
}

// PrintUnderline formats the given string with ANSI underline code and outputs it to the console
//
// Param str: string to be formatted
func PrintUnderline(str string) {
	print(underline + str + resetUnderline)
}

// Println formats the given string with ANSI color codes and outputs it to the console, appending it with the "\n" character
//
// Param str: string to be formatted
//
// Param colorFg: foreground color of the text
//
// Param colorBg: background color of the text
//
// Note: Even if your console does not support RGB color codes, it will still select the closest color from its palette
func Println(str string, colorFg color.Color, colorBg color.Color) {
	println(makeBackgroundColorEscapeSequence(colorBg) + makeForegroundColorEscapeSequence(colorFg) + str + reset)
}

// PrintlnFg formats the given string with ANSI foreground color code and outputs it to the console, appending it with the "\n" character
//
// Param str: string to be formatted
//
// Param color: foreground color of the text
//
// Note: Even if your console does not support RGB color codes, it will still select the closest color from its palette
func PrintlnFg(str string, color color.Color) {
	println(makeForegroundColorEscapeSequence(color) + str + reset)
}

// PrintlnBg formats the given string with ANSI background color code and outputs it to the console, appending it with the "\n" character
//
// Param str: string to be formatted
//
// Param color: background color of the text
//
// Note: Even if your console does not support RGB color codes, it will still select the closest color from its palette
func PrintlnBg(str string, color color.Color) {
	println(makeBackgroundColorEscapeSequence(color) + str + reset)
}

// PrintlnBold formats the given string with ANSI bold code and outputs it to the console, appending it with the "\n" character
//
// Param str: string to be formatted
func PrintlnBold(str string) {
	println(bold + str + resetBold)
}

// PrintlnUnderline formats the given string with ANSI underline code and outputs it to the console, appending it with the "\n" character
//
// Param str: string to be formatted
func PrintlnUnderline(str string) {
	println(underline + str + resetUnderline)
}

// Sprint formats the given string with ANSI color codes and returns it
//
// Param str: string to be formatted
//
// Param colorFg: foreground color of the text
//
// Param colorBg: background color of the text
//
// Note: Even if your console does not support RGB color codes, it will still select the closest color from its palette
func Sprint(str string, colorFg color.Color, colorBg color.Color) string {
	return makeBackgroundColorEscapeSequence(colorBg) + makeForegroundColorEscapeSequence(colorFg) + str + reset
}

// SprintFg formats the given string with ANSI foreground color code and returns it
//
// Param str: string to be formatted
//
// Param color: foreground color of the text
//
// Note: Even if your console does not support RGB color codes, it will still select the closest color from its palette
func SprintFg(str string, color color.Color) string {
	return makeForegroundColorEscapeSequence(color) + str + reset
}

// SprintBg formats the given string with ANSI background color code and returns it
//
// Param str: string to be formatted
//
// Param color: background color of the text
//
// Note: Even if your console does not support RGB color codes, it will still select the closest color from its palette
func SprintBg(str string, color color.Color) string {
	return makeBackgroundColorEscapeSequence(color) + str + reset
}

// SprintBold formats the given string with ANSI bold code and returns it
//
// Param str: string to be formatted
func SprintBold(str string) string {
	return bold + str + resetBold
}

// SprintUnderline formats the given string with ANSI underline code and returns it
//
// Param str: string to be formatted
func SprintUnderline(str string) string {
	return underline + str + resetUnderline
}

// Setting the color and style separately

// SetForegroundColor sets the foreground color of the console output to given color until the next Reset call
//
// Param color: foreground color to be set
func SetForegroundColor(color color.Color) {
	print(makeForegroundColorEscapeSequence(color))
}

// SetBackgroundColor sets the background color of the console output to given color until the next Reset call
//
// Param color: background color to be set
func SetBackgroundColor(color color.Color) {
	print(makeBackgroundColorEscapeSequence(color))
}

// SetBold sets the console output to be in bold style until the next ResetBold or Reset call
func SetBold() {
	print(bold)
}

// ResetBold resets the bold style
func ResetBold() {
	print(resetBold)
}

// SetUnderline sets the console output to be in underline style until the next ResetUnderline or Reset call
func SetUnderline() {
	print(underline)
}

// ResetUnderline resets the underline style
func ResetUnderline() {
	print(resetUnderline)
}

// Reset resets foreground, background colors and text styles
func Reset() {
	print(reset)
}
