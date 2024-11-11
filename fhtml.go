package fhtml

type (
	// Renderer is interface for views and layouts
	Renderer interface {
		Render(data []byte) Renderer
		Layout() Renderer
		SetLayout(layout Renderer) Renderer
		Builder() *Builder
	}
)
