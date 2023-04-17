package fhtml

import "bytes"

type (
	Comment struct {
		Text  string
		Nodes []Node
	}
)

func Cmt(c string, nodes ...Node) *Comment {
	return &Comment{
		Text:  c,
		Nodes: nodes,
	}
}

func (c *Comment) Render(bb *bytes.Buffer, n int) {
	newLineIndent(bb, n)
	writeStrings(bb, `<!-- `, c.Text)

	if len(c.Nodes) > 0 {
		for _, e := range c.Nodes {
			e.Render(bb, n+1)
		}
		newLineIndent(bb, n)
	}

	bb.WriteString(` -->`)
}
