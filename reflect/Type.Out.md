# Out(i int) Type

参数列表

- i int 函数输出参数的索引号

返回值：

- Type 函数输出参数的 reflect.Type 类型

功能说明：

- reflect.TypeOf(interface{}).Out(int) 返回函数输出的参数类型。如果出现恐慌（panic），表示该类型不是Func或索引号超出范围。

代码实例：
  
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a = func() (int, string, uint, float64) {
			return 1, "1", 2, 1.2
		}
		var typeof reflect.Type = reflect.TypeOf(a)
		
		for i:=0; i<typeof.NumOut(); i++ {
			fmt.Println(typeof.Out(i).Kind())
			// 0 >>int
			// 1 >>string
			// 2 >>uint
			// 3 >>float64
		}
	}
