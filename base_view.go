package fhtml

type (
	BaseView struct {
		builder *Builder
		layout  Renderer
	}
)

func (v *BaseView) Run() []byte {
	layout := Renderer(v)
	for layout.Layout() != nil {
		layout = layout.Layout()
	}

	return layout.Render().Bytes()
}

func (v *BaseView) Render() Renderer {
	return v
}

func (v *BaseView) Layout() Renderer {
	return v.layout
}

func (v *BaseView) Builder() *Builder {
	return v.builder
}

func (v *BaseView) Bytes() []byte {
	return v.builder.Bytes()
}

func (v *BaseView) String() string {
	return v.builder.String()
}
