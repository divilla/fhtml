package fhtml

type (
	BaseComponent struct {
		builder *Builder
	}
)

func (v *BaseComponent) Builder() *Builder {
	return v.builder
}

func (v *BaseComponent) SetBuilder(builder *Builder) {
	v.builder = builder
}

func (v *BaseComponent) Bytes() []byte {
	return v.builder.Bytes()
}

func (v *BaseComponent) String() string {
	return v.builder.String()
}
