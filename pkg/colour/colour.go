package colour

import (
	"fmt"
	"strconv"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/termenv"
)

// Get the current terminal colour profile
func Profile() termenv.Profile {
	return termenv.ColorProfile()
}

// Color a string's foreground with the given value.
func Foreground(val, color string) string {
	return termenv.String(val).Foreground(Profile().Color(color)).String()
}

// Return a function that will colorize the foreground of a given string.
func MakeFgStyle(color string) func(string) string {
	return termenv.Style{}.Foreground(Profile().Color(color)).Styled
}

// Color a string's foreground and background with the given value.
func MakeFgBgStyle(fg, bg string) func(string) string {
	return termenv.
		Style{}.
		Foreground(Profile().Color(fg)).
		Background(Profile().Color(bg)).
		Styled
}

// Generate a blend of colors.
func MakeRamp(colorA, colorB string, steps float64) (s []string) {
	cA, _ := colorful.Hex(colorA)
	cB, _ := colorful.Hex(colorB)

	for i := 0.0; i < steps; i++ {
		c := cA.BlendLuv(cB, i/steps)
		s = append(s, ToHex(c))
	}
	return
}

// Convert a colorful.Color to a hexadecimal format compatible with termenv.
func ToHex(c colorful.Color) string {
	return fmt.Sprintf("#%s%s%s", FloatToHex(c.R), FloatToHex(c.G), FloatToHex(c.B))
}

// Helper function for converting colors to hex. Assumes a value between 0 and
// 1.
func FloatToHex(f float64) (s string) {
	s = strconv.FormatInt(int64(f*255), 16)
	if len(s) == 1 {
		s = "0" + s
	}
	return
}
