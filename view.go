package fhtml

import (
	"github.com/tidwall/gjson"
)

type (
	View struct {
		builder *Builder
		layout  Renderer
	}
)

func NewView(data []byte) *View {
	v := &View{
		builder: NewBuilder(data),
	}
	v.layout = NewInnerLayout(v)

	return v
}

func (v *View) Render() Renderer {
	b := v.builder
	b.EC(`<section class="section">`).C(
		b.GetForeach(`nums`, func(key, val gjson.Result) {
			b.EC(`<div>`).C(
				b.E(`<h1 class="title">`, `Hello World `, val.Raw, `</h1>`),
				b.GetIf(`show`, func() {
					b.EC(`<p class="subtitle">`).C(
						b.E(`My first website with <strong>Bulma</strong>!`),
					).E(`</p>`)
				}),
			).E(`</div>`)
		}),
	).E(`</section>`)

	return v
}

func (v *View) Layout() Renderer {
	return v.layout
}

func (v *View) Builder() *Builder {
	return v.builder
}

func (v *View) Bytes() []byte {
	return v.builder.Bytes()
}

func (v *View) String() string {
	return v.builder.String()
}
