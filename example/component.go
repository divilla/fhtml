package main

import (
	"github.com/divilla/fhtml"
)

type (
	Component struct {
		fhtml.BaseComponent
	}
)

func NewComponent(builder *fhtml.Builder) *Component {
	v := new(Component)
	v.SetBuilder(builder)

	return v
}

func (v *Component) Render(nr string) bool {
	b := v.Builder()
	b.E(`<h1 class="title">`, `Hello World `, nr, `</h1>`)

	return true
}
