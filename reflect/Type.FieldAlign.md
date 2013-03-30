# FieldAlign() int

参数列表

- 无

返回值：

- int 字段对齐的值

功能说明：

- reflect.TypeOf(interface{}).FieldAlign() 返回字段对齐的值（以字节为单位）

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	type A struct {
		A0 int "标签0"
		A1 string "标签1"
		A2 []byte "标签2"
	}
	
	func main(){
		var a A
		var typeof reflect.Type = reflect.TypeOf(a.A0)
		fmt.Println(typeof.FieldAlign())
		//>>4
		typeof = reflect.TypeOf(a.A1)
		fmt.Println(typeof.FieldAlign())
		//>>4
		typeof = reflect.TypeOf(a.A2)
		fmt.Println(typeof.FieldAlign())
		//>>4
	}
