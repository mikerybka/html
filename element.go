package html

import (
	"bytes"
	"io"
	"net/http"
)

// Element represents an HTML DOM node.
// Text elements are defined by an empty tag and a "string" attribute.
type Element struct {
	Tag      string
	Attrs    Attrs
	Children []Element
}

// String returns the HTML representation of the element.
func (e Element) String() string {
	buf := bytes.NewBuffer(nil)
	e.Write(buf)
	return buf.String()
}

// ServeHTTP implements the http.Handler interface.
func (e Element) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	e.Write(w)
}

func (e Element) Write(w io.Writer) (int, error) {
	// Handle text nodes
	if e.Tag == "" {
		text := e.Attrs["string"]
		if text == "" {
			return 0, nil
		}
		return w.Write([]byte(text))
	}

	// Write element tag
	written, err := w.Write([]byte("<" + e.Tag))
	if err != nil {
		return written, err
	}

	// Write element attributes if any
	for k, v := range e.Attrs {
		n, err := w.Write([]byte(" " + k + "=\"" + v + "\""))
		if err != nil {
			return written + n, err
		}
		written += n
	}

	// Write element children if any
	if len(e.Children) == 0 {
		n, err := w.Write([]byte(" />"))
		if err != nil {
			return written + n, err
		}
		return written + n, nil
	}
	n, err := w.Write([]byte(">"))
	if err != nil {
		return written + n, err
	}
	written += n
	for _, child := range e.Children {
		n, err := child.Write(w)
		if err != nil {
			return written + n, err
		}
		written += n
	}

	// Write closing tag
	n, err = w.Write([]byte("</" + e.Tag + ">"))
	if err != nil {
		return written + n, err
	}
	return written + n, nil
}
