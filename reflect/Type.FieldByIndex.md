# FieldByIndex(index []int) StructField

参数列表

- index []int 传入字段与嵌套字段的索引号

返回值：

- .StructField 返回字段的 reflect.StructField 类型

功能说明：

- reflect.TypeOf(interface{}).FieldByIndex([]int) 返回对应的索引序列的嵌套字段。如果出现恐慌（panic），因为type不是struct类型。

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
		B "匿名嵌套"
	}
	type B struct {
		B0 int "B结构-标签B0"
	}
	
	func main(){
		var a A
		var typeof reflect.Type = reflect.TypeOf(a)
		
		var r = []int{3,4}
		var f = typeof.FieldByIndex(r)
		fmt.Println(f.Index, f.PkgPath, f.Name, f.Type, f.Tag, f.Anonymous)
		//>>[4]  B main.B 匿名嵌套 true
		
		r = []int{4,0}
		f = typeof.FieldByIndex(r)
		fmt.Println(f.Index, f.PkgPath, f.Name, f.Type, f.Tag, f.Anonymous)
		//>>[0]  B0 int B结构-标签B0 false
	}
