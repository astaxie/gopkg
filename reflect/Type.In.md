# In(int) Type

参数列表

- int 函数参数的索引号

返回值：

- Type 函数参数的 reflect.Type 类型

功能说明：

- reflect.TypeOf(interface{}).In(int) 函数输入的参数类型，如果出现恐慌（panic），该类型不是函数或索引号超出范围。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a = func(a1 int, a2 string, a3 uint, a4 interface{}) {}
		var typeof reflect.Type = reflect.TypeOf(a)
		for i:=0; i<typeof.NumIn(); i++ {
			fmt.Println(typeof.In(i).Kind())
			// 0 >>int
			// 1 >>string
			// 2 >>uint
			// 3 >>interface
		}
	}
