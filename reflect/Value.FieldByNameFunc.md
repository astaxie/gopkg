# func (v Value) FieldByNameFunc(match func(string) bool) Value

参数列表

- match func(string) bool 使用函数功能来筛选想要的字段。

返回值：

- Value  返回字段的 reflect.Vaue 类型

功能说明：

- reflect.ValueOf(interface{}).FieldByNameFunc(func(string) bool) 传入字段“字符串”名称，并判断，func 返回 true，返回字段的 Value 类型。如果出现恐慌(panic)，表示该类型不是struct。

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
		var f = func(s string) bool {
			if s == "A1" { // 结合正则匹配，功能最强大。
				return true
			}
			return false
		}
		var a A
		var value reflect.Value = reflect.ValueOf(a)
		vf := value.FieldByNameFunc(f)
		fmt.Println(vf.Kind(), vf.Uint(), vf.Kind() == reflect.Uint)
		//>>uint 0 true
	}
