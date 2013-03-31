func (enc *Encoder) Encode(v interface{}) error

参数列表:

- v 序列化json对象

返回值:

- error 错误

功能说明:

这个函数主要是将encode编码的JSON数据写入到相关联的对象

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