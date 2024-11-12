package main

import "github.com/divilla/fhtml"

type InnerLayout struct {
	fhtml.BaseLayout
}

func NewInnerLayout(content fhtml.Renderer) *InnerLayout {
	l := new(InnerLayout)
	l.SetContent(content)
	l.SetLayout(NewOuterLayout(l))

	return l
}

func (l *InnerLayout) Render(b *fhtml.Builder, data []byte) *fhtml.Builder {
	b.EC(`<div class="container">`).C(
		l.Content().Render(b, data),
	).E(`</div>`)

	return b
}
