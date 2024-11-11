package fhtml

import (
	"github.com/tidwall/gjson"
)

type (
	View struct {
		BaseView
	}
)

func NewView(data []byte) *View {
	v := new(View)
	v.builder = NewBuilder(data)
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
