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

	layout.Render(b, data)

	return b.Bytes()
}

func (v *BaseView) Render(b *Builder, data []byte) *struct{} {
	_ = b
	_ = data
	return nil
}

func (v *BaseView) Layout() Renderer {
	return v.layout
}

func (v *BaseView) SetLayout(layout Renderer) Renderer {
	v.layout = layout
	return v
}
