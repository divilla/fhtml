package fhtml

import (
	"bytes"
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

// HTML writes indented raw strings - not sanitized HTML
func (b *Builder) HTML(tokens ...string) *Builder {
	b.indent()
	return b.HTMLInline(tokens...)
}

// HTMLInline writes raw strings inline - not sanitized HTML
func (b *Builder) HTMLInline(tokens ...string) *Builder {
	b.WriteString(tokens...)
	return b
}

// Text writes indented sanitized text
func (b *Builder) Text(text string) *Builder {
	b.indent()
	return b.TextInline(text)
}

// TextInline writes sanitized text inline
func (b *Builder) TextInline(text string) *Builder {
	ss, err := htmlsanitizer.SanitizeString(text)
	if err != nil {
		panic(err)
	}
	b.bb.WriteString(ss)

	return b
}

// Attr writes HTML attribute
func (b *Builder) Attr(vals ...string) string {
	if len(vals) == 0 || vals[0] == "" {
		return zeroString
	}
	if len(vals) == 1 {
		return vals[0]
	}

	return ` ` + vals[0] + `="` + vals[1] + `"`
}

// Class writes HTML class attribute
func (b *Builder) Class(vals ...string) string {
	res := ` class="`
	for key, val := range vals {
		if key > 0 {
			res += " "
		}
		res += val
	}
	res += `"`

	return res
}

// Tag builds HTML element
func (b *Builder) Tag(tag string, attrs ...string) *Builder {
	b.indent()
	b.pushTag(tag)

	return b.TagInline(tag, attrs...)
}

// TagInline builds inline HTML element
func (b *Builder) TagInline(tag string, attrs ...string) *Builder {
	b.WriteString(`<` + tag)
	b.WriteString(attrs...)
	b.WriteString(`>`)

	return b
}

// Children builds child elements
func (b *Builder) Children(tags ...any) *Builder {
	_ = tags

	tag := b.popTag()
	b.indent()
	b.bb.WriteString(`</` + tag + `>`)

	return b
}

// ChildrenInline builds inline child elements
func (b *Builder) ChildrenInline(tags ...any) *Builder {
	_ = tags

	tag := b.popTag()
	b.bb.WriteString(`</` + tag + `>`)

	return b
}

// Void prevents closing void tab
func (b *Builder) Void() *Builder {
	b.popTag()
	return b
}

// Document creates base HTML document
func (b *Builder) Document(tags ...any) *Builder {
	_ = tags
	return b
}

// If executes 'fn' if result of 'expression' is true
func (b *Builder) If(expression bool, s string) string {
	if expression {
		return s
	}

	return ""
}

// IfFunc executes 'fn' if result of 'expression' is true
func (b *Builder) IfFunc(expression bool, fn func()) *struct{} {
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

func (b *Builder) pushTag(tag string) {
	b.tags = append(b.tags, tag)
}

func (b *Builder) popTag() string {
	l := len(b.tags) - 1
	tag := b.tags[l]
	b.tags = b.tags[:l]

	return tag
}

func (b *Builder) indent() {
	indent := len(b.tags)

	if val, ok := indentCache[indent]; ok {
		b.bb.WriteString(val)
		return
	}

	val := "\n"
	for j := 0; j < indent; j++ {
		val += Indent
	}
	indentCache[indent] = val
	b.bb.WriteString(val)
}
