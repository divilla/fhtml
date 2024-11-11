package fhtml

import (
	"bytes"
	"strconv"

	"github.com/tidwall/gjson"
)

var (
	Indent      = `  `
	indentCache = make(map[int]string)
)

type (
	Builder struct {
		data []byte
		bb   *bytes.Buffer
		ind  int
	}

	Renderer interface {
		Render() Renderer
		Layout() Renderer
		Builder() *Builder
		Bytes() []byte
		String() string
	}
)

// NewBuilder constructs *Builder provided 'data' argument is valid JSON
func NewBuilder(data []byte) *Builder {
	return &Builder{
		data: data,
		bb:   new(bytes.Buffer),
	}
}

// H writes raw strings - not sanitized HTML
func (b *Builder) H(tokens ...string) *Builder {
	for _, token := range tokens {
		b.bb.WriteString(token)
	}

	return b
}

// E is used for writing elements without Children
func (b *Builder) E(tokens ...string) *Builder {
	indent(b.bb, b.ind)
	for _, token := range tokens {
		b.bb.WriteString(token)
	}

	return b
}

// EC is used for writing elements with Children
func (b *Builder) EC(tokens ...string) *Builder {
	indent(b.bb, b.ind)
	for _, token := range tokens {
		b.bb.WriteString(token)
	}
	b.ind++

	return b
}

// C is building Element's Children
func (b *Builder) C(a ...any) *Builder {
	b.ind--

	return b
}

// D creates base HTML document
func (b *Builder) D(a ...any) *Builder {
	return b
}

// GetString extracts string value from JSON data for provided path
func (b *Builder) GetString(path string) string {
	return gjson.GetBytes(b.data, path).Raw
}

// GetIf extracts bool value from JSON data for provided 'path' and executes 'fn' if result is true
func (b *Builder) GetIf(path string, fn func()) *Builder {
	if gjson.GetBytes(b.data, path).Bool() {
		fn()
	}

	return b
}

// GetForeach extracts array from JSON data for provided 'path' and executes 'fn' for each array member
func (b *Builder) GetForeach(path string, fn func(key, val gjson.Result)) *Builder {
	for key, val := range gjson.GetBytes(b.data, path).Array() {
		fn(keyResult(key), val)
	}

	return b
}

// Bytes returns []byte form Buffer
func (b *Builder) Bytes() []byte {
	return b.bb.Bytes()
}

// String returns string form Buffer
func (b *Builder) String() string {
	return b.bb.String()
}

func indent(bb *bytes.Buffer, ind int) {
	if val, ok := indentCache[ind]; ok {
		bb.WriteString(val)
		return
	}

	s := "\n"
	for j := 0; j < ind; j++ {
		s += Indent
	}
	indentCache[ind] = s

	bb.WriteString(s)
}

func keyResult(key int) gjson.Result {
	return gjson.Result{
		Type: gjson.Number,
		Raw:  strconv.Itoa(key),
		Str:  strconv.Itoa(key),
		Num:  float64(key),
	}
}
