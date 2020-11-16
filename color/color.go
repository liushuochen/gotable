package color

import (
	"fmt"
)

type Color string

var (
	GREENBG      = Color([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	WHITEBG      = Color([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	YELLOWBG     = Color([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	REDBG        = Color([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	BLUEBG       = Color([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	MAGENTABG    = Color([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	CYANBG       = Color([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	GREEN        = Color([]byte{27, 91, 51, 50, 109})
	WHITE        = Color([]byte{27, 91, 51, 55, 109})
	YELLOW       = Color([]byte{27, 91, 51, 51, 109})
	RED          = Color([]byte{27, 91, 51, 49, 109})
	BLUE         = Color([]byte{27, 91, 51, 52, 109})
	MAGENTA      = Color([]byte{27, 91, 51, 53, 109})
	CYAN         = Color([]byte{27, 91, 51, 54, 109})
	RESET        = Color([]byte{27, 91, 48, 109})
	DISABLECOLOR = ""
)

func PrintlnGreen(obj interface{}) {
	fmt.Println(GREEN, obj, RESET)
}

func PrintlnRed(obj interface{}) {
	fmt.Println(GREEN, obj, RESET)
}

type String struct {
	Color    Color
	Origin   string
	Colorful string
	Length   int
}

func (c String) Val() string {
	return c.Colorful
}

func (c String) Len() int {
	return c.Length
}

func ColorfulString(color Color, obj interface{}) String {
	origin := fmt.Sprintf("%v", obj)

	if len(color) == 0 {
		return String{
			Color:    color,
			Colorful: origin,
			Length:   len(origin),
		}
	}

	return String{
		Color:    color,
		Colorful: fmt.Sprintf("%v%v%v", color, obj, RESET),
		Length:   len(origin),
	}
}
