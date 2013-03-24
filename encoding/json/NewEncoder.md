func NewEncoder(w io.Writer) *Encoder

参数列表:

- r Writer对象

返回值:

- *Encoderr 指向Encoder的指针

功能说明:

这个函数主要是给w创建一个encoder实例

代码实例:

    package main

	import (
		"encoding/json"
		"fmt"
		"os"
	)
	
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	
	func main() {
		group := ColorGroup{
			ID:     1,
			Name:   "Reds",
			Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		}
		encoder := json.NewEncoder(os.Stdout)
		if err := encoder.Encode(group); err != nil {
			fmt.Printf("failed encoding to writer: %s", err)
		}
	}






显示结果:

	{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}