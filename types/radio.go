package pearls

import (
	"fmt"

	"github.com/minibase-app/pearls/pkg/colour"
)

type Radio struct {
	// Selected bool
	Checked bool
	Label   string
}

func (radio *Radio) Select() {
	radio.Checked = true
}

func (radio *Radio) Deselect() {
	radio.Checked = false
}

func (radio *Radio) Toggle() {
	radio.Checked = !radio.Checked
}

func (radio *Radio) Render() string {
	var fill rune
	var foreground string

	if radio.Checked {
		fill = '•'
		// fill = '⚈' // IDEA: alt?
		foreground = "#017fb0"
	} else {
		fill = ' '
		foreground = " "
	}

	return colour.Foreground(
		fmt.Sprintf("(%s) %s", string(fill), radio.Label),
		foreground,
	)
}
