package fhtml

import (
	"bytes"
)

type (
	Attribute struct {
		Name   string
		Value  string
		IsBool bool
	}
)

func Attr(s ...string) *Attribute {
	if len(s) > 1 {
		return &Attribute{
			Name:  s[0],
			Value: s[1],
		}
	}

	return &Attribute{
		Name:   s[0],
		IsBool: true,
	}
}

func (a *Attribute) Add(ss ...string) *Attribute {
	for _, v := range ss {
		a.Value += " " + v
	}

	return a
}

func (a *Attribute) AddIf(b bool, ss ...string) *Attribute {
	if b {
		a.Add(ss...)
	}

	return a
}

func (a *Attribute) Render(bb *bytes.Buffer, n int) {
	bb.WriteString(` `)
	bb.WriteString(a.Name)

	if a.IsBool {
		return
	}

	bb.WriteString(`="`)
	bb.WriteString(a.Value)
	bb.WriteString(`"`)
}
