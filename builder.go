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
		bba  *bytes.Buffer
		tags []string
	}

	BuilderFn func(b *Builder) *struct{}
)

// NewBuilder constructs *Builder provided 'data' argument is valid JSON
func NewBuilder() *Builder {
	return &Builder{
		bb:  new(bytes.Buffer),
		bba: new(bytes.Buffer),
	}
}

// HI writes raw strings inline - not sanitized HTML
func (b *Builder) HI(tokens ...string) *struct{} {
	for _, token := range tokens {
		b.bb.WriteString(token)
	}

	return nil
}

// HTML writes indented raw strings - not sanitized HTML
func (b *Builder) HTML(tokens ...string) *struct{} {
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
func (b *Builder) A(attr ...string) *struct{} {
	if len(attr) == 1 && attr[0] != "" {
		return b.WriteStringAfter(" ", attr[0])
	}
	if len(attr) > 1 && attr[0] != "" {
		return b.WriteStringAfter(" ", attr[0], `="`, attr[1], `"`)
	}

	return nil
}

// Class writes class attribute
func (b *Builder) Class(vals ...string) *struct{} {
	b.WriteStringAfter(` class="`)
	for key, val := range vals {
		if val == "" {
			continue
		}
		if key > 0 {
			b.WriteStringAfter(" ")
		}
		b.WriteStringAfter(val)
	}
	return b.WriteStringAfter(`""`)
}

// TagInline is used for writing elements without Children
func (b *Builder) TagInline(tag string, attrs ...any) *Builder {
	_ = attrs

	b.WriteString(`<`, tag)
	b.bb.Write(b.bba.Bytes())
	b.WriteString(`>`)
	b.tags = append(b.tags, tag)
	b.bba.Reset()

	return b
}

// TagVoid is used for writing void elements
func (b *Builder) TagVoid(tag string, attrs ...any) *struct{} {
	_ = attrs

	indent(b.bb, len(b.tags))
	b.TagInline(tag)
	popTag(b)

	return nil
}

// Tag is used for writing elements
func (b *Builder) Tag(tag string, attrs ...any) *Builder {
	_ = attrs

	indent(b.bb, len(b.tags))
	b.TagInline(tag)

	return b
}

// ContentInline is building Element's Children inline
func (b *Builder) ContentInline(a ...any) *struct{} {
	_ = a

	tag := popTag(b)

	return b.WriteString(`</`, tag, `>`)
}

// Content is building Element's Children
func (b *Builder) Content(a ...any) *struct{} {
	_ = a

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

// WriteStringAfter returns string form Buffer
func (b *Builder) WriteStringAfter(s ...string) *struct{} {
	for _, v := range s {
		_, err := b.bba.WriteString(v)
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
