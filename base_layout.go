package fhtml

type (
	BaseLayout struct {
		layout  Renderer
		content Renderer
	}
)

func (v *BaseLayout) Render(b *Builder, data []byte) *Builder {
	_ = data
	return b
}

func (l *BaseLayout) Content() Renderer {
	return l.content
}

func (l *BaseLayout) SetContent(content Renderer) Renderer {
	l.content = content
	return l
}

func (l *BaseLayout) Layout() Renderer {
	return l.layout
}

func (l *BaseLayout) SetLayout(layout Renderer) Renderer {
	l.layout = layout
	return l
}
