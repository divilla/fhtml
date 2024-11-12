package fhtml

type (
	BaseView struct {
		layout Renderer
	}
)

func (v *BaseView) Run(b *Builder, data []byte) []byte {
	layout := Renderer(v)
	for layout.Layout() != nil {
		layout = layout.Layout()
	}

	return layout.Render(b, data).Bytes()
}

func (v *BaseView) Render(b *Builder, data []byte) *Builder {
	_ = data
	return b
}

func (v *BaseView) Layout() Renderer {
	return v.layout
}

func (v *BaseView) SetLayout(layout Renderer) Renderer {
	v.layout = layout
	return v
}
