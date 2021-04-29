package pearls

import "strings"

type Multiselect struct {
	Items    []Checkbox
	Size     int
	Selected int
}

// Display the selection list
func (list *Multiselect) Render() string {
	var items []string

	for _, element := range list.Items {
		items = append(items, element.Render())
	}

	return strings.Join(items, "\n")
}

// Go back to the previous selection in the list from the current location
func (list *Multiselect) Previous() {
	list.Select(list.Selected - 1)
}

// Fast forward the selection all the way to the end.
func (list *Multiselect) End() {
	list.Select(list.Size)
}

// Advance the selection position.
func (list *Multiselect) Next() {
	list.Select(list.Selected + 1)
}

// Rewind the selection all the way to the start.
func (list *Multiselect) Rewind() {
	list.Select(0)
}

func (list *Multiselect) Select(index int) (*Checkbox, *Checkbox) {
	previous := &list.Items[list.Selected]

	previous.Deselect()

	if index < 0 {
		index = 0
	} else if index > (list.Size - 1) {
		index = list.Size - 1
	}

	item := &list.Items[index]
	list.Selected = index

	item.Select()

	return item, previous
}
