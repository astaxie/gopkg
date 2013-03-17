# type StructField

参数列表

- 无

返回值：

- .Index []int 返回当前的字段索引号
- .PkgPath string 返回当前字段的包路径
- .Name string 返回当前的字段名称
- .Offect uintptr 返回当前的字段偏移量，从结构中的字段开始到当前字段开始的偏移量
- .Type reflect.Type 返回当前的字段 Type 类型
- .Tag reflect.StructTag 返回当前的字段标签名称
- .Anonymous bool 返回当前的字段是否是匿名字段

功能说明：

- StructField 在一个结构内描述了一个单一的字段。

代码实例：

	package main
	import (
	    "fmt"
	    "reflect"
	)

	type A struct {
		A0 int "这是A0"
		A1 B "嵌套B结构"
		B "匿名字段"
	}
	type B struct {
		B0 string
		B1 int
	}
	func main(){
		var a A
		var typeof reflect.Type = reflect.TypeOf(a)
		var field reflect.StructField = typeof.Field(0) //0 表示是 a 结构中的第几位字段
		fmt.Println(field.Index, field.PkgPath, field.Name, field.Offset, field.Type, field.Tag, field.Anonymous)
		//>>[0]  A0 0 int 这是A0 false

		var field1 reflect.StructField = typeof.Field(1)
		fmt.Println(field1.Index, field1.PkgPath, field1.Name, field1.Offset, field1.Type, field1.Tag, field1.Anonymous)
		//>>[1]  A1 4 main.B 嵌套B结构 false

		var field2 reflect.StructField = typeof.Field(2)
		fmt.Println(field2.Index, field2.PkgPath, field2.Name, field2.Offset, field2.Type, field2.Tag, field2.Anonymous)
		//>>[2]  B 16 main.B 匿名字段 true
	}
