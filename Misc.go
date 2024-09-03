package GoColor

import (
	"strconv"
)

// Miscellaneous stuff, like cursor position

func ResetCursor() {
	print(escSequence + "H")
}

func MoveCursor(x int, y int) {
	print(escSequence + strconv.Itoa(y) + strconv.Itoa(x) + "H")
}

func Clear() {
	print(escSequence + "2J")
}
