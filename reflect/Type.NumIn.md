# NumIn() int

参数列表

- 无

返回值：

- int 函数传入参数的总数量

功能说明：

- reflect.TypeOf(interface{}).NumIn() 函数传入参数的总数量，如果出现恐慌(panic)，该类型不是 Func。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a = func(a1 int, a2 string, a3 uint, a4 interface{}) {}
		var typeof reflect.Type = reflect.TypeOf(a)
		fmt.Println(typeof.NumIn())
		//>>4
	}
