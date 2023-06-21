# html

`html` is a Go package that provides a set of types, interfaces and functions for creating UI components directly in your Go source code.

## Core Concepts

The package provides an `Element` type that represents an HTML DOM node:
```go
type Element struct {
	Tag      string
	Attrs    Attrs
	Children []Element
}
```

as well as a `Component` interface:
```go
type Component interface {
	Render() Element
}
```
