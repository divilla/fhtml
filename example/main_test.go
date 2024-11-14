package main

import (
	"testing"

	"github.com/divilla/fhtml"

	"github.com/tidwall/sjson"
)

var (
	result []byte
)

func Benchmark_Render(b *testing.B) {
	var data, r []byte
	data, _ = sjson.SetBytes(data, `title`, `Hello Bulma!`)
	data, _ = sjson.SetBytes(data, `nums`, []int{1, 2, 3, 4, 5, 6})
	data, _ = sjson.SetBytes(data, `show`, true)

	for n := 0; n < b.N; n++ {
		view := NewView(data)
		builder := fhtml.NewBuilder()
		r = view.Run(builder)
		builder.Close()
	}

	result = r
}
