package fhtml

type OuterLayout struct {
	*BaseLayout
}

func NewOuterLayout(content Renderer) *OuterLayout {
	return &OuterLayout{
		&BaseLayout{
			builder: content.Builder(),
			content: content,
		},
	}
}

func (l *OuterLayout) Render() Renderer {
	b := l.builder
	b.D(
		b.H(`<!DOCTYPE html>`),
		b.EC(`<html>`).C(
			b.EC(`<head>`).C(
				b.E(`<meta charset="utf-8">`),
				b.E(`<meta name="viewport" content="width=device-width, initial-scale=1">`),
				b.E(`<title>`, b.GetString(`title`), `</title>`),
				b.E(`<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@1.0.2/css/bulma.min.css">`),
			).E(`</head>`),
			b.EC(`<body>`).C(
				l.content.Render(),
			).E(`</body>`),
		).E(`</html>`),
	)

	return l
}
