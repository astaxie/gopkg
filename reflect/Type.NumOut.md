# NumOut() int

参数列表

- 无

返回值：

- int 函数输出参数总数量

功能说明：

- reflect.TypeOf(interface{}).NumOut() 一个函数的输出参数总数量。如果出现恐慌（panic），表示该类型不是Func。

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
		fmt.Println(typeof.NumOut())
		//>>4
	}
