package pearls

import (
	"fmt"
	"math"
	"strings"

	"github.com/minibase-app/pearls/pkg/colour"
	"github.com/muesli/termenv"
)

const (
	progressBarWidth  = 71
	progressFullChar  = "█"
	progressEmptyChar = "░"
)

// General stuff for styling the view
var (
	subtle        = colour.MakeFgStyle("241")
	progressEmpty = subtle(progressEmptyChar)
	dot           = colour.Foreground(" • ", "236")

	// Gradient colors we'll use for the progress bar
	ramp = colour.MakeRamp("#01a96e", "#017fb0", progressBarWidth)
)

type Progress struct {
	Percent float64
}

func (bar *Progress) Advance(new float64) {
	if new >= 1 {
		bar.Percent = 1
	}

	bar.Percent = new // TODO: make this additive instead of setting
}

// Report the current state of the progress bar
func (bar *Progress) Current() float64 {
	return bar.Percent
}

func (bar *Progress) Render() string {
	w := float64(progressBarWidth)

	fullSize := int(math.Round(w * bar.Percent))
	var fullCells string
	for i := 0; i < fullSize; i++ {
		fullCells += termenv.
			String(progressFullChar).
			Foreground(colour.Profile().Color(ramp[i])).
			String()
	}

	emptySize := int(w) - fullSize
	emptyCells := strings.Repeat(progressEmpty, emptySize)

	return fmt.Sprintf("%s%s %3.0f", fullCells, emptyCells, math.Round(bar.Percent*100))
}
