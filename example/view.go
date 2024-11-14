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
	return b.Tag("section", b.Class("section")).Content(
		b.Foreach(data, `nums`, func(key, val gjson.Result) {
			b.Tag("div").Content(
				b.Tag("h1", b.Class("title")).ContentInline(
					b.HI("Hello World ", val.Raw),
				),
				NewComponent().Render(b, val.Raw),
				b.If(b.GetBool(data, `show`), func() {
					b.Tag("p", b.Class("subtitle")).Content(
						b.HTML(`My first website with <strong>Bulma</strong>!`),
					)
				}),
			)
		}),
	)
}
