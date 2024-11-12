package fhtml

import (
	"bytes"
	"strings"

	"github.com/tidwall/sjson"

	"github.com/sym01/htmlsanitizer"

	"github.com/tidwall/gjson"
)

var (
	Indent      = `  `
	indentCache = make(map[int]string)
)

type (
	Builder struct {
		bb  *bytes.Buffer
		ind int
	}
)

// NewBuilder constructs *Builder provided 'data' argument is valid JSON
func NewBuilder() *Builder {
	return &Builder{
		bb: new(bytes.Buffer),
	}
}

// H writes raw strings - not sanitized HTML
func (b *Builder) H(tokens ...string) *struct{} {
	for _, token := range tokens {
		b.bb.WriteString(token)
	}

	return nil
}

// HI writes indented raw strings - not sanitized HTML
func (b *Builder) HI(tokens ...string) *struct{} {
	indent(b.bb, b.ind)
	return b.H(tokens...)
}

// T writes sanitized text
func (b *Builder) T(tokens ...string) *struct{} {
	s := strings.Join(tokens, ``)
	ss, err := htmlsanitizer.SanitizeString(s)
	if err != nil {
		panic(err)
	}
	b.bb.WriteString(ss)

	return nil
}

// TI writes indented sanitized strings
func (b *Builder) TI(tokens ...string) *struct{} {
	indent(b.bb, b.ind)
	return b.T(tokens...)
}

// E is used for writing elements without Children
func (b *Builder) E(tokens ...string) *struct{} {
	indent(b.bb, b.ind)
	for _, token := range tokens {
		b.bb.WriteString(token)
	}

	return nil
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
	_ = a
	b.ind--

	return b
}

// D creates base HTML document
func (b *Builder) D(a ...any) *struct{} {
	_ = a
	return nil
}

// If executes 'fn' if result of 'expression' is true
func (b *Builder) If(expression bool, fn func()) *struct{} {
	if expression {
		fn()
	}

	return nil
}

// Foreach extracts array from JSON data for provided 'path' and executes 'fn' for each array member
func (b *Builder) Foreach(data []byte, path string, fn func(key, value gjson.Result)) *struct{} {
	gjson.GetBytes(data, path).ForEach(func(key, value gjson.Result) bool {
		fn(key, value)
		return true
	})

	return nil
}

// GetString extracts string value from JSON data given provided path
func (b *Builder) GetString(data []byte, path string) string {
	return gjson.GetBytes(data, path).Raw
}

// GetBool extracts bool value from JSON data given provided path
func (b *Builder) GetBool(data []byte, path string) bool {
	return gjson.GetBytes(data, path).Bool()
}

// SetData builds JSON
func (b *Builder) SetData(data []byte, path string, value interface{}) []byte {
	o, _ := sjson.SetBytes(data, path, value)
	return o
}

// Bytes returns []byte form Buffer
func (b *Builder) Bytes() []byte {
	return b.bb.Bytes()
}

// String returns string form Buffer
func (b *Builder) String() string {
	return b.bb.String()
}

// Close destroys Buffer
func (b *Builder) Close() {
	b.bb.Reset()
	b.bb = nil
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
