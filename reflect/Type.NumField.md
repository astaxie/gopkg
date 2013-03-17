# NumField() int

参数列表

- 无

返回值：

- int 字段总数量

功能说明：

- reflect.TypeOf(interface{}).NumField() 返回struct结构中的字段总数量

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	type A struct {
		A0 int
		A1 string
		A2 []byte
	}

	func main(){
		var a A
		var typeof reflect.Type = reflect.TypeOf(a)
		fmt.Println(typeof.NumField())
		//>>3
	}
