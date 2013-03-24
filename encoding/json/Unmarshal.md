func Unmarshal(data []byte, v interface{}) error

参数列表:

- data json数据
- v 反序列化JSON对象

返回值:

- error 错误

功能说明:

这个函数主要是用于反序列化JSON编码的data，并将结果存储到指向v的指针

代码实例:

    package main

	import (
		"encoding/json"
		"fmt"
	)
	
	func main() {
		var jsonBlob = []byte(`[
			{"Name": "Platypus", "Order": "Monotremata"},
			{"Name": "Quoll",    "Order": "Dasyuromorphia"}
		]`)
		type Animal struct {
			Name  string
			Order string
		}
		var animals []Animal
		err := json.Unmarshal(jsonBlob, &animals)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("%+v", animals)
	}




显示结果:

	[{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]