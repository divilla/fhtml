package fhtml

type (
	BaseView struct {
		data   []byte
		layout Renderer
	}
)

func (v *BaseView) Run(b *Builder) {
	layout := Renderer(v)
	for layout.Layout() != nil {
		layout = layout.Layout()
	}

	layout.Render(b)
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

func (v *BaseView) Render(b *Builder) *Builder {
	_ = b
	return nil
}
