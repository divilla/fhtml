package main

import (
	"github.com/divilla/fhtml"
)

type (
	Component struct {
		fhtml.BaseComponent
	}
)

func NewComponent() *Component {
	v := new(Component)
	return v
}

func (v *Component) Render(b *fhtml.Builder, nr string) *struct{} {
	return b.Tag("h1", b.Class("test", "title")).ContentInline(b.HTMLInline("Hello World ", nr))
}
