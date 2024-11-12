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

func (v *Component) Render(b *fhtml.Builder, nr string) *fhtml.Builder {
	b.E(`<h1 class="title">`, `Hello World `, nr, `</h1>`)

	return b
}
