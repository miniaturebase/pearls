package radio

import (
	"strings"

	pearls "github.com/minibase-app/pearls/types"
)

func New(label string, checked bool) pearls.Radio {
	return pearls.Radio{Checked: checked, Label: label}
}

func List(elements ...pearls.Radio) string {
	var list []string

	for _, element := range elements {
		list = append(list, element.Render())
	}

	return strings.Join(list, "\n")
}
