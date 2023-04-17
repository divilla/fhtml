package fhtml

import (
	"bytes"
)

type Element struct {
	Tag        string
	Attributes []*Attribute
	Elements   []*Element
	InnerHTML  string
	IsVoid     bool
}

func Elm(t string, as ...*Attribute) *Element {
	return &Element{
		Tag:        t,
		Attributes: as,
	}
}

func (e *Element) Content(s string) {
	e.InnerHTML = s
}

func (e *Element) Children(es ...*Element) {
	e.Elements = es
}

func (e *Element) Void() {
	e.IsVoid = true
}

func (e *Element) Render(bb *bytes.Buffer, n int) {
	newLine(bb)
	indent(bb, n)

	bb.WriteString(`<`)
	bb.WriteString(e.Tag)

	for _, attr := range e.Attributes {
		attr.Render(bb, 0)
	}

	bb.WriteString(`>`)

	if e.IsVoid {
		return
	}

	bb.WriteString(e.InnerHTML)

	for _, elm := range e.Elements {
		elm.Render(bb, n+1)
	}

	bb.WriteString(`<\`)
	bb.WriteString(e.Tag)
	bb.WriteString(`>`)
}

var IndentSize = 4
var indentsBytes = makeIndents(256)
var newLineByte = []byte("\n")

func makeIndents(n int) [][]byte {
	var indent, indentString string
	indents := make([][]byte, n)

	for i := 0; i < IndentSize; i++ {
		indentString += ` `
	}

	for i := 0; i < n; i++ {
		indent = ``
		for j := 0; j < i; j++ {
			indent += indentString
		}
		indents[i] = []byte(indent)
	}

	return indents
}

func indent(bb *bytes.Buffer, n int) {
	bb.Write(indentsBytes[n])
}

func newLine(bb *bytes.Buffer) {
	bb.Write(newLineByte)
}

func startCommentIf(bb *bytes.Buffer, cond bool) {
	if cond {
		bb.WriteString(`<!-- `)
	}
}

func endCommentIf(bb *bytes.Buffer, cond bool) {
	if cond {
		bb.WriteString(` -->`)
	}
}
