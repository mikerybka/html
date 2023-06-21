package html

// Components are types with a Render method that returns an Element.
// They are used as building blocks for user interfaces.
type Component interface {
	Render() Element
}
