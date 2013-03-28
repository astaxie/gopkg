# FieldByName(name string) (StructField, bool)

参数列表

- name string 传入一个“字符串”的字段名称

返回值：

- .StructField 返回字段的 reflect.StructField 类型
- bool 如果找到字段，返回true，则返回false

功能说明：

- reflect.TypeOf(interface{}).FieldByName(string) 使用一个“字符串”的字段名称返回字段的 StructField 类型和布尔值。布尔值为true，表示找到该字段。

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
		A3 *A "指针指向自身"
	}
	
	func main(){
		var a A
		var typeof reflect.Type = reflect.TypeOf(a)
		
		var f, ok = typeof.FieldByName("A2")
		fmt.Println(ok, f.Index, f.PkgPath, f.Name, f.Type, f.Tag, f.Anonymous)
		//>>true [2]  A2 []uint8 标签2 false
	}
