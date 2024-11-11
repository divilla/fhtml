package main

import (
	"testing"

	"github.com/tidwall/sjson"
)

var result []byte

func Benchmark_Render(b *testing.B) {
	var j, r []byte
	j, _ = sjson.SetBytes(j, `title`, `Hello Bulma!`)
	j, _ = sjson.SetBytes(j, `nums`, []int{1, 2, 3, 4, 5, 6})
	j, _ = sjson.SetBytes(j, `show`, true)

	for n := 0; n < b.N; n++ {
		r = NewView(j).Run()
	}

	result = r
}
