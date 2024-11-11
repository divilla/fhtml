package fhtml

func FindOutermostLayout(layout Renderer) Renderer {
	if layout == nil {
		return nil
	}

	for layout.Layout() != nil {
		layout = layout.Layout()
	}

	return layout
}

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

func (l *BaseLayout) Layout() Renderer {
	return l.layout
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
