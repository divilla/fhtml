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

func (l *OuterLayout) Render(b *fhtml.Builder) *fhtml.Builder {
	return b.Document(
		b.HTMLInline(`<!DOCTYPE html>`),
		b.Tag("html").Children(
			b.Tag("head").Children(
				b.Tag("meta", b.Attr("charset", "utf-8")).Void(),
				b.Tag("meta", b.Attr("name", "viewport"), b.Attr("content", "width=device-width, initial-scale=1")).Void(),
				b.Tag("title").ChildrenInline(
					b.HTMLInline(l.title()),
				),
				b.Tag("link", b.Attr("rel", "stylesheet"), b.Attr("href", "https://cdn.jsdelivr.net/npm/bulma@1.0.2/css/bulma.min.css")).Void(),
			),
			b.Tag("body").Children(
				l.Content().Render(b),
			),
		),
	)
}

func (l *OuterLayout) title() string {
	return gjson.GetBytes(l.Data(), "title").Raw
}
