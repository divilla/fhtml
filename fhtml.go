package fhtml

type (
	// Renderer is interface for views and layouts
	Renderer interface {
		Layout() Renderer
		SetLayout(layout Renderer) Renderer
		Render(b *Builder) *struct{}
	}
)

// Foreach iterates through slice executing callback for each member
func Foreach[T any](data []T, fn func(key int, value T) bool) *struct{} {
	for key, value := range data {
		if fn(key, value) == false {
			break
		}
	}

	return nil
}
