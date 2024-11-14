package main

import (
	"github.com/divilla/fhtml"
	"github.com/tidwall/gjson"
)

type OuterLayout struct {
	fhtml.BaseLayout
}

func NewOuterLayout(content fhtml.Renderer) *OuterLayout {
	l := new(OuterLayout)
	l.SetContent(content)
	l.SetData(content.Data())

	return l
}

func (l *OuterLayout) Render(b *fhtml.Builder) *struct{} {
	return b.Document(
		b.HTMLInline(`<!DOCTYPE html>`),
		b.Tag("html").Content(
			b.Tag("head").Content(
				b.TagVoid("meta", b.A("charset", "utf-8")),
				b.TagVoid("meta", b.A("name", "viewport").A("content", "width=device-width, initial-scale=1")),
				b.Tag("title").ContentInline(b.HTMLInline(l.title())),
				b.TagVoid("link", b.A("rel", "stylesheet").A("href", "https://cdn.jsdelivr.net/npm/bulma@1.0.2/css/bulma.min.css")),
			),
			b.Tag("body").Content(
				l.Content().Render(b),
			),
		),
	)
}

func (l *OuterLayout) title() string {
	return gjson.GetBytes(l.Data(), "title").Raw
}
