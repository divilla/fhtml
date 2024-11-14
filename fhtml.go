package fhtml

type (
	// Renderer is interface for views and layouts
	Renderer interface {
		Data() []byte
		SetData(data []byte) Renderer
		Layout() Renderer
		SetLayout(layout Renderer) Renderer
		Render(b *Builder) *struct{}
	}
)
