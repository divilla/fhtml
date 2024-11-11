package fhtml

type (
	// Renderer is interface for layouts, views and components
	Renderer interface {
		Render() Renderer
		Layout() Renderer
		Builder() *Builder
		Bytes() []byte
		String() string
	}
)

func FindOutermostLayout(layout Renderer) Renderer {
	if layout == nil {
		return nil
	}

	for layout.Layout() != nil {
		layout = layout.Layout()
	}

	return layout
}
