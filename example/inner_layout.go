package main

import "github.com/divilla/fhtml"

type InnerLayout struct {
	fhtml.BaseLayout
}

func NewInnerLayout(content fhtml.Renderer) *InnerLayout {
	l := new(InnerLayout)
	l.SetContent(content)
	l.SetData(content.Data())
	l.SetLayout(NewOuterLayout(l))

	return l
}

func (l *InnerLayout) Render(b *fhtml.Builder) *struct{} {
	return b.Tag("div", b.Class("container")).Content(
		l.Content().Render(b),
	)
}
