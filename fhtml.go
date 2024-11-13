package fhtml

type (
	// Renderer is interface for views and layouts
	Renderer interface {
		Render(b *Builder, data []byte) *struct{}
		Layout() Renderer
		SetLayout(layout Renderer) Renderer
	}
)
