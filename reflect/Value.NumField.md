# func (v Value) NumField() int

参数列表

- 无

返回值：

- int  返回字段的数量

功能说明：

- reflect.ValueOf(interface{}).NumField() 返回struct结构的字段数量。如果出现恐慌(panic)，表示该类型不是struct。

代码实例：
  
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	type A struct {
		A0 int
		A1 string
		A2 uint
		A4 float64
		A5 interface{}
	}
	
	func main(){
		var a A
		var value reflect.Value = reflect.ValueOf(a)
		fmt.Println(value.NumField())
		//>>5
	}
