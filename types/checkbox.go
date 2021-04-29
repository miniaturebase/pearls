package pearls

import (
	"fmt"

	"github.com/minibase-app/pearls/pkg/colour"
)

const (
	template string = "[%s] %s"
	checked  rune   = '✔'
	empty    rune   = ' '
)

// The checkbox element and it's state.
type Checkbox struct {
	// Is the checkbox currently checked?
	Checked bool

	// The string value rendered to the screen for users.
	Label string
}

// Render the checkbox in it's entirety.
func (checkbox *Checkbox) Render() string {
	var fill rune
	var foreground string

	if checkbox.Checked {
		fill = checked
		foreground = "#017fb0"
	} else {
		fill = empty
		foreground = ""
	}

	return colour.Foreground(
		fmt.Sprintf(template, string(fill), checkbox.Label),
		foreground,
	)
}

// Select the checkbox.
func (checkbox *Checkbox) Select() {
	checkbox.Checked = true
}

// Deselect a selected checkbox.
func (checkbox *Checkbox) Deselect() {
	checkbox.Checked = false
}

// Toggle the checkbox state.
func (checkbox *Checkbox) Toggle() {
	checkbox.Checked = !checkbox.Checked
}
