# func (v Value) FieldByIndex(index []int) Value

参数列表

- index []int 嵌套字段的索引号

返回值：

- Value  指定字段的返回 reflect.Vaue 类型

功能说明：

- reflect.ValueOf(interface{}).FieldByIndex([]int) 对应索引号返回嵌套字段的 Value 类型。如果出现恐慌(panic)，表示该类型不是struct。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	type A struct {
		A0 int
		A1 B
	}
	type B struct {
		B0 string
	}
	
	func main(){
		var a A
		var value reflect.Value = reflect.ValueOf(a)
		var tt = []int{1,0} //1 表示 A1 B，0 表示 B0 string
		vf := value.FieldByIndex(tt)
		fmt.Println(vf.Kind(), vf.String(), vf.Kind() == reflect.String)
		//>>string  true
	}
