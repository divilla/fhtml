package main

import "github.com/divilla/fhtml"

type OuterLayout struct {
	fhtml.BaseLayout
}

func NewOuterLayout(content fhtml.Renderer) *OuterLayout {
	l := new(OuterLayout)
	l.SetContent(content)
	return l
}

func (l *OuterLayout) Render(b *fhtml.Builder, data []byte) *struct{} {
	return b.D(
		b.HTML(`<!DOCTYPE html>`),
		b.Tag("html").Content(
			b.Tag("head").Content(
				b.TagVoid("meta", b.A("charset", "utf-8")),
				b.TagVoid("meta", b.A("name", "viewport"), b.A("content", "width=device-width, initial-scale=1")),
				b.Tag("title").ContentInline(b.HI(b.GetString(data, "title"))),
				b.TagVoid("link", b.A("rel", "stylesheet"), b.A("href", "https://cdn.jsdelivr.net/npm/bulma@1.0.2/css/bulma.min.css")),
			),
			b.Tag("body").Content(
				l.Content().Render(b, data),
			),
		),
	)
}
