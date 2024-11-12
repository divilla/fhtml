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

func (l *OuterLayout) Render(b *fhtml.Builder, data []byte) *fhtml.Builder {
	b.D(
		b.H(`<!DOCTYPE html>`),
		b.EC(`<html>`).C(
			b.EC(`<head>`).C(
				b.E(`<meta charset="utf-8">`),
				b.E(`<meta name="viewport" content="width=device-width, initial-scale=1">`),
				b.E(`<title>`, b.GetString(data, `title`), `</title>`),
				b.E(`<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@1.0.2/css/bulma.min.css">`),
			).E(`</head>`),
			b.EC(`<body>`).C(
				l.Content().Render(b, data),
			).E(`</body>`),
		).E(`</html>`),
	)

	return b
}
