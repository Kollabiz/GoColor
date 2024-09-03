package GoColor

const (
	Black          = 0
	DarkRed        = 1
	DarkGreen      = 2
	DarkYellow     = 3
	DarkBlue       = 4
	DarkMagenta    = 5
	DarkTurquoise  = 6
	Gray           = 7
	DarkGray       = 8
	LightRed       = 9
	LightGreen     = 10
	LightYellow    = 11
	LightBlue      = 12
	LightMagenta   = 13
	LightTurquoise = 14
	White          = 15
)

func Grayscale(brightness int) int {
	if brightness > 15 || brightness < 0 {
		panic("grayscale color must be in range of 0-15")
	}
	return brightness + 232
}
