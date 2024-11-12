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
	v.SetLayout(NewInnerLayout(v))

	return v
}

func (v *View) Render(b *fhtml.Builder, data []byte) *fhtml.Builder {
	b.EC(`<section class="section">`).C(
		b.Foreach(data, `nums`, func(key, val gjson.Result) {
			b.EC(`<div>`).C(
				b.E(`<h1 class="`, b.IfString(true, `title`), `">`, `Hello World `, val.Raw, `</h1>`),
				NewComponent().Render(b, val.Raw),
				b.If(b.GetBool(data, `show`), func() {
					b.EC(`<p class="subtitle">`).C(
						b.E(`My first website with <strong>Bulma</strong>!`),
					).E(`</p>`)
				}),
			).E(`</div>`)
		}),
	).E(`</section>`)

	return b
}
