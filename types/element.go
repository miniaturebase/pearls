package pearls

// "✔"
// "✖"
// "⚈"
// "•"
// "→"
// "!"
// "?"
// "∅"
// "⌽"

type Element interface {
	Render() string
}

type Selectable interface {
	Select()
	Deselect()
}

type Toggleable interface {
	Toggle()
}
