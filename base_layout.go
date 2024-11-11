package fhtml

type (
	BaseLayout struct {
		builder *Builder
		layout  Renderer
		content Renderer
	}
)

func (l *BaseLayout) Render() Renderer {

	return l
}

func (l *BaseLayout) Content() Renderer {
	return l.content
}

func (l *BaseLayout) SetContent(content Renderer) Renderer {
	l.content = content
	l.builder = content.Builder()
	return l
}

func (l *BaseLayout) Layout() Renderer {
	return l.layout
}

func (l *BaseLayout) SetLayout(layout Renderer) Renderer {
	l.layout = layout
	return l
}

func (l *BaseLayout) Builder() *Builder {

	return l.builder
}

func (l *BaseLayout) Bytes() []byte {
	return l.builder.Bytes()
}

func (l *BaseLayout) String() string {
	return l.builder.String()
}
