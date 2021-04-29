package checkbox

import pearls "github.com/minibase-app/pearls/types"

// Create a new checkbox element with the given label and checked state.
func New(label string, checked bool) pearls.Checkbox {
	return pearls.Checkbox{
		Checked: checked,
		Label:   label,
	}
}

// Create a new checkbox element in the checked state, with the given lavel.
func Checked(label string) pearls.Checkbox {
	return New(label, true)
}

// Create a new checkbox element in the unchecked state, with the given lavel.
func Unchecked(label string) pearls.Checkbox {
	return New(label, false)
}

// Create a list of checkboxes.
func List(labels ...string) pearls.Multiselect {
	var length int = len(labels)
	var elements = make([]pearls.Checkbox, length)

	for index, label := range labels {
		elements[index] = Unchecked(label)
	}

	return pearls.Multiselect{
		Items:    elements,
		Size:     length,
		Selected: 0,
	}
}
