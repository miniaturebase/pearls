package pearls

import "strings"

type Select struct {
	Items    []Radio
	Size     int
	Selected int
}

// Display the selection list
func (list *Select) Render() string {
	var items []string

	for _, element := range list.Items {
		items = append(items, element.Render())
	}

	return strings.Join(items, "\n")
}

// Go back to the previous selection in the list from the current location
func (list *Select) Previous() {
	list.Select(list.Selected - 1)
}

// Advance the selection position.
func (list *Select) Next() {
	list.Select(list.Selected + 1)
}

// Rewind the selection all the way to the start.
func (list *Select) Rewind() {
	list.Select(0)
}

// Fast forward the selection all the way to the end.
func (list *Select) End() {
	list.Select(list.Size)
}

func (list *Select) Select(index int) Radio {
	if index > (list.Size - 1) {
		index = list.Size - 1
	}

	item := list.Items[index]
	previous := list.Selected
	list.Selected = index

	list.Items[previous].Deselect()
	item.Select()

	return item
}
