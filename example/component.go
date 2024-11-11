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

func (v *Component) Render(path string) bool {
	b := v.Builder()
	b.E(`<h1 class="title">`, `Hello World `, b.GetString(path), `</h1>`)
	return true
}
