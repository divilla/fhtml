package main

import (
	"fmt"

	"github.com/divilla/fhtml"

	"github.com/tidwall/sjson"
)

func main() {
	var j []byte
	j, _ = sjson.SetBytes(j, `title`, `Hello Bulma!`)
	j, _ = sjson.SetBytes(j, `nums`, []int{1, 2, 3, 4, 5, 6})
	j, _ = sjson.SetBytes(j, `show`, true)
	v := NewView(j)

	b := fhtml.NewBuilder()
	v.Run(b)
	fmt.Println(string(b.Bytes()))
	b.Close()
}
