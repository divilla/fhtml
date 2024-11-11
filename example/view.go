package main

import (
	"github.com/divilla/fhtml"
	"github.com/tidwall/gjson"
)

type (
	View struct {
		fhtml.BaseView
	}
)

func NewView() *View {
	v := new(View)
	v.SetBuilder(fhtml.NewBuilder())
	v.SetLayout(NewInnerLayout(v))

	return v
}

func (v *View) Render(data []byte) fhtml.Renderer {
	b := v.Builder()
	b.EC(`<section class="section">`).C(
		b.Foreach(data, `nums`, func(key, val gjson.Result) {
			b.EC(`<div>`).C(
				b.E(`<h1 class="title">`, `Hello World `, val.Raw, `</h1>`),
				NewComponent(b).Render(data, `nums.`+key.Raw),
				b.If(b.GetBool(data, `show`), func() {
					b.EC(`<p class="subtitle">`).C(
						b.E(`My first website with <strong>Bulma</strong>!`),
					).E(`</p>`)
				}),
			).E(`</div>`)
		}),
	).E(`</section>`)

	return v
}
