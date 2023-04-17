package fhtml

import "bytes"

type (
	Node interface {
		Render(bb *bytes.Buffer, n int)
	}
)
