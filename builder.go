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

// HTML writes indented raw strings - not sanitized HTML
func (b *Builder) HTML(tokens ...string) *struct{} {
	indent(b.bb, len(b.tags))
	return b.HTMLInline(tokens...)
}

// HTMLInline writes raw strings inline - not sanitized HTML
func (b *Builder) HTMLInline(tokens ...string) *struct{} {
	for _, token := range tokens {
		b.bb.WriteString(token)
	}

	return nil
}

// Text writes indented sanitized text
func (b *Builder) Text(tokens ...string) *struct{} {
	indent(b.bb, len(b.tags))
	return b.TextInline(tokens...)
}

// TextInline writes sanitized text inline
func (b *Builder) TextInline(tokens ...string) *struct{} {
	s := strings.Join(tokens, ``)
	ss, err := htmlsanitizer.SanitizeString(s)
	if err != nil {
		panic(err)
	}
	b.bb.WriteString(ss)

	return nil
}

// A writes HTML attribute
func (b *Builder) A(attr ...string) *Builder {
	if len(attr) == 1 && attr[0] != "" {
		b.WriteStringAfter(" ", attr[0])
	}
	if len(attr) > 1 && attr[0] != "" {
		b.WriteStringAfter(" ", attr[0], `="`, attr[1], `"`)
	}

	return b
}

// Class writes HTML class attribute
func (b *Builder) Class(vals ...string) *Builder {
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
	b.WriteStringAfter(`""`)

	return b
}

// Tag is used for writing elements
func (b *Builder) Tag(tag string, attrs ...any) *Builder {
	_ = attrs

	indent(b.bb, len(b.tags))
	b.TagInline(tag)

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

// Content is building Element's Children
func (b *Builder) Content(a ...any) *struct{} {
	_ = a

	tag := popTag(b)
	indent(b.bb, len(b.tags))

	return b.WriteString(`</`, tag, `>`)
}

// ContentInline is building Element's Children inline
func (b *Builder) ContentInline(a ...any) *struct{} {
	_ = a

	tag := popTag(b)

	return b.WriteString(`</`, tag, `>`)
}

// Inline writes raw string content into Tag
func (b *Builder) Inline(tokens ...string) *struct{} {
	for _, token := range tokens {
		b.bb.WriteString(token)
	}

	tag := popTag(b)

	return b.WriteString(`</`, tag, `>`)
}

// Document creates base HTML document
func (b *Builder) Document(a ...any) *struct{} {
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
