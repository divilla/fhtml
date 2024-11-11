package fhtml

type InnerLayout struct {
	*BaseLayout
}

func NewInnerLayout(content Renderer) *InnerLayout {
	l := &InnerLayout{
		&BaseLayout{
			builder: content.Builder(),
			content: content,
		},
	}
	l.layout = NewOuterLayout(l)

	return l
}

func (l *InnerLayout) Render() Renderer {
	b := l.builder
	b.EC(`<div class="container">`).C(
		l.content.Render(),
	).E(`</div>`)

	return l
}
