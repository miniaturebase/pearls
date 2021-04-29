package pearls

import (
	"fmt"
	"strings"

	"github.com/minibase-app/pearls/pkg/colour"
)

type Menu struct {
	Items    []string
	Selected int
}

func (list *Menu) Render() string {
	items := make([]string, len(list.Items))
	var highlight string

	for index, item := range list.Items {
		if list.Selected == index {
			highlight = "#107fb0"
		} else {
			highlight = ""
		}

		items[index] = fmt.Sprintf(colour.Foreground("%s", highlight), item)
	}

	return strings.Join(items, "\n")
}

// Go back to the previous selection in the list from the current location
func (list *Menu) Previous() {
	list.Select(list.Selected - 1)
}

// Advance the selection position.
func (list *Menu) Next() {
	list.Select(list.Selected + 1)
}

// Rewind the selection all the way to the start.
func (list *Menu) Rewind() {
	list.Select(0)
}

// Fast forward the selection all the way to the end.
func (list *Menu) End() {
	list.Select(list.Size())
}

func (list *Menu) Size() int {
	return len(list.Items)
}

func (list *Menu) Select(index int) {
	length := list.Size()

	if index < 0 {
		index = 0
	} else if index > (length - 1) {
		index = length - 1
	}

	list.Selected = index
}
