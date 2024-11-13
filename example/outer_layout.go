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
		b.H(`<!DOCTYPE html>`),
		b.E("html").C(
			b.E("head").C(
				b.EV("meta", b.A("charset", "utf-8")),
				b.EV("meta", b.A("name", "viewport"), b.A("content", "width=device-width, initial-scale=1")),
				b.E("title").CI(b.HI(b.GetString(data, "title"))),
				b.EV("link", b.A("rel", "stylesheet"), b.A("href", "https://cdn.jsdelivr.net/npm/bulma@1.0.2/css/bulma.min.css")),
			),
			b.E("body").C(
				l.Content().Render(b, data),
			),
		),
	)
}
