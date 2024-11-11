package main

import (
	"fmt"

	"github.com/tidwall/sjson"
)

func main() {
	var j []byte
	j, _ = sjson.SetBytes(j, `title`, `Hello Bulma!`)
	j, _ = sjson.SetBytes(j, `nums`, []int{1, 2, 3, 4, 5, 6})
	j, _ = sjson.SetBytes(j, `show`, true)
	fmt.Println(string(j))
	v := NewView().Run(j)
	fmt.Println(string(v))
}
