# Field(i int) StructField

参数列表

- i int 传入字段的索引号

返回值：

- .StructField 返回字段的 reflect.StructField 类型

功能说明：

- reflect.TypeOf(interface{}).Field(int) 返回一个结构中的指定字段类型。

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
		B "匿名嵌套"
	}
	type B struct {}
	
	func main(){
		var a A
		var typeof reflect.Type = reflect.TypeOf(a)
		for i:=0; i<typeof.NumField(); i++ {
			f := typeof.Field(i)
			fmt.Println(f.Index, f.PkgPath, f.Name, f.Type, f.Tag, f.Anonymous)
			// 0 >>[0]  A0 int 标签0 false
			// 1 >>[1]  A1 string 标签1 false
			// 2 >>[2]  A2 []uint8 标签2 false
			// 3 >>[3]  B main.B 匿名嵌套 true
		}
	}
