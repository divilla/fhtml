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

func NewView(data []byte) *View {
	v := new(View)
	v.SetData(data)
	v.SetLayout(NewInnerLayout(v))

	return v
}

func (v *View) Render(b *fhtml.Builder) *fhtml.Builder {
	return b.Tag("section", b.Class("section")).Children(
		b.Foreach(v.Data(), `nums`, func(key, val gjson.Result) {
			b.Tag("div").Children(
				b.Tag("h1", b.Class("title"), b.Attr("id", "")).ChildrenInline(
					b.HTMLInline("Hello World ", val.Raw),
				),
				NewComponent().Render(b, val.Raw),
				b.IfFunc(b.GetBool(v.Data(), `show`), func() {
					b.Tag("p", b.Class("subtitle")).Children(
						b.HTML(`My first website with <strong>Bulma</strong>!`),
					)
				}),
			)
		}),
	)
}
