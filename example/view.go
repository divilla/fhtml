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

func (v *View) Render(b *fhtml.Builder, data []byte) *struct{} {
	return b.E("section", b.Class("section")).C(
		b.Foreach(data, `nums`, func(key, val gjson.Result) {
			b.E("div").C(
				b.E("h1", b.Class("title")).CI(b.HI("Hello World ", val.Raw)),
				NewComponent().Render(b, val.Raw),
				b.If(b.GetBool(data, `show`), func() {
					b.E("p", b.Class("subtitle")).C(
						b.H(`My first website with <strong>Bulma</strong>!`),
					)
				}),
			)
		}),
	)
}
