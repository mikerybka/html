package html

// Text is a convenience function for creating a text node.
func Text(s string) Element {
	return Element{Tag: "", Attrs: Attrs{"string": s}}
}
