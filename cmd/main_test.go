package main

import (
	"testing"

	"github.com/divilla/fhtml"
	"github.com/tidwall/sjson"
)

var result []byte

func Benchmark_Render(b *testing.B) {
	var j, r []byte
	j, _ = sjson.SetBytes(j, `title`, `Hello Bulma!`)
	j, _ = sjson.SetBytes(j, `nums`, []int{1, 2, 3, 4, 5, 6})
	j, _ = sjson.SetBytes(j, `show`, true)

	for n := 0; n < b.N; n++ {
		v := fhtml.NewView(j)
		r = fhtml.FindOutermostLayout(v).Render().Bytes()
	}

	result = r
}
