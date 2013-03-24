package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func main() {
	dst := new(bytes.Buffer)
	src := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	json.Indent(dst, src, "##", "**")
	json.Indent(dst, src, "##", "**")
	fmt.Println(dst)
}
