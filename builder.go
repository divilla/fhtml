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
		bb   *bytes.Buffer
		tags []string
	}

	BuilderFn func(b *Builder) *struct{}
)

// NewBuilder constructs *Builder provided 'data' argument is valid JSON
func NewBuilder() *Builder {
	return &Builder{
		bb: new(bytes.Buffer),
	}
}

// HI writes raw strings inline - not sanitized HTML
func (b *Builder) HI(tokens ...string) *struct{} {
	for _, token := range tokens {
		b.bb.WriteString(token)
	}

	return nil
}

// H writes indented raw strings - not sanitized HTML
func (b *Builder) H(tokens ...string) *struct{} {
	indent(b.bb, len(b.tags))
	return b.HI(tokens...)
}

// TI writes sanitized text inline
func (b *Builder) TI(tokens ...string) *struct{} {
	s := strings.Join(tokens, ``)
	ss, err := htmlsanitizer.SanitizeString(s)
	if err != nil {
		panic(err)
	}
	b.bb.WriteString(ss)

	return nil
}

// T writes indented sanitized strings
func (b *Builder) T(tokens ...string) *struct{} {
	indent(b.bb, len(b.tags))
	return b.TI(tokens...)
}

// A writes attribute
func (b *Builder) A(attr ...string) BuilderFn {
	return func(b *Builder) *struct{} {
		if len(attr) == 1 && attr[0] != "" {
			return b.WriteString(` `, attr[0])
		}
		if len(attr) > 1 && attr[0] != "" {
			return b.WriteString(` `, attr[0], `="`, attr[1], `"`)
		}

		return nil
	}
}

// Class writes class attribute
func (b *Builder) Class(vals ...string) BuilderFn {
	return func(b *Builder) *struct{} {
		b.WriteString(` class="`)
		for key, val := range vals {
			if val == "" {
				continue
			}
			if key > 0 {
				b.WriteString(` `)
			}
			b.WriteString(val)
		}
		return b.WriteString(`"`)
	}
}

// EI is used for writing elements without Children
func (b *Builder) EI(tag string, attrs ...BuilderFn) *Builder {
	b.WriteString(`<`, tag)
	for _, attr := range attrs {
		attr(b)
	}
	b.WriteString(`>`)
	b.tags = append(b.tags, tag)

	return b
}

// EV is used for writing void elements
func (b *Builder) EV(tag string, attrs ...BuilderFn) *struct{} {
	indent(b.bb, len(b.tags))
	b.EI(tag, attrs...)
	popTag(b)

	return nil
}

// E is used for writing elements
func (b *Builder) E(tag string, attrs ...BuilderFn) *Builder {
	indent(b.bb, len(b.tags))
	b.EI(tag, attrs...)

	return b
}

// CI is building Element's Children inline
func (b *Builder) CI(a ...any) *struct{} {
	_ = a

	if len(b.tags) == 0 {
		return nil
	}

	tag := popTag(b)

	return b.WriteString(`</`, tag, `>`)
}

// C is building Element's Children
func (b *Builder) C(a ...any) *struct{} {
	_ = a

	if len(b.tags) == 0 {
		return nil
	}

	tag := popTag(b)
	indent(b.bb, len(b.tags))

	return b.WriteString(`</`, tag, `>`)
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

// IfString executes 'fn' if result of 'expression' is true
func (b *Builder) IfString(expression bool, s string) string {
	if expression {
		return s
	}

	return ""
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

// WriteString returns string form Buffer
func (b *Builder) WriteString(s ...string) *struct{} {
	for _, v := range s {
		_, err := b.bb.WriteString(v)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

// Close destroys Buffer
func (b *Builder) Close() {
	b.bb.Reset()
	b.bb = nil
}

func popTag(b *Builder) string {
	l := len(b.tags) - 1
	tag := b.tags[l]
	b.tags = b.tags[:l]

	return tag
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
