# func (v Value) FieldByName(name string) Value

参数列表

- name string 指定一个“字符串”的字段名称

返回值：

- Value  指定字段的返回 reflect.Vaue 类型

功能说明：

- reflect.ValueOf(interface{}).FieldByName(string) 指定一个“字符串”字段名称返回字段的 Value 类型。如果出现恐慌(panic)，表示该类型不是struct。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	type A struct {
		A0 int
		A1 uint
	}
	func main(){
		var a A
		var value reflect.Value = reflect.ValueOf(a)
		vf := value.FieldByName("A1")
		fmt.Println(vf.Kind(), vf.Uint(), vf.Kind() == reflect.Uint)
		//>>uint 0 true
	}
