package fhtml

type (
	BaseLayout struct {
		layout  Renderer
		content Renderer
		data    []byte
	}
)

func (l *BaseLayout) Data() []byte {
	return l.data
}

func (l *BaseLayout) SetData(data []byte) Renderer {
	l.data = data
	return l
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

func (l *BaseLayout) Render(b *Builder) *struct{} {
	_ = b
	return nil
}
