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

func (v *BaseView) SetLayout(layout Renderer) Renderer {
	v.layout = layout
	return v
}

func (v *BaseView) Builder() *Builder {
	return v.builder
}

func (v *BaseView) SetBuilder(builder *Builder) Renderer {
	v.builder = builder
	return v
}

func (v *BaseView) Bytes() []byte {
	return v.builder.Bytes()
}

func (v *BaseView) String() string {
	return v.builder.String()
}
