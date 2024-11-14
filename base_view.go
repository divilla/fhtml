package fhtml

type (
	BaseView struct {
		layout Renderer
		data   []byte
	}
)

func (v *BaseView) Run(b *Builder) []byte {
	layout := Renderer(v)
	for layout.Layout() != nil {
		layout = layout.Layout()
	}

	layout.Render(b)

	return b.Bytes()
}

func (v *BaseView) Data() []byte {
	return v.data
}

func (v *BaseView) SetData(data []byte) Renderer {
	v.data = data
	return v
}

func (v *BaseView) Layout() Renderer {
	return v.layout
}

func (v *BaseView) SetLayout(layout Renderer) Renderer {
	v.layout = layout
	return v
}

func (v *BaseView) Render(b *Builder) *struct{} {
	_ = b
	return nil
}
