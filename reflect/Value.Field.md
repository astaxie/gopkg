# func (v Value) Field(i int) Value

参数列表

- i int 字段的索引号

返回值：

- Value  指定字段的返回 reflect.Vaue 类型

功能说明：

- reflect.ValueOf(interface{}).Field(int) 指定索引号返回struct中字段的 Value 类型。如果出现恐慌(panic)，表示该类型不是struct。

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
		for i:=0; i<value.NumField(); i++ {
			vf := value.Field(i)
			fmt.Println(vf.Kind(), vf.Interface()) //Value 的方法有N多，这里就不一一演示。后面将一一介绍
			// 0 >>int 0
			// 1 >>string 
			// 2 >>uint 0
			// 3 >>float64 0
			// 4 >>interface <nil>
		}
	}
