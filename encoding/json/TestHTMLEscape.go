package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	dst := new(bytes.Buffer)
	src := []byte(`{
		"Name":"tony.shao",
		"Age":25,
		"Job":"Programmer<Escaping>"
		}`)
	json.HTMLEscape(dst, src)
	fmt.Println(dst)
}
