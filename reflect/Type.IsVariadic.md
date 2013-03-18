# IsVariadic() bool

参数列表

- 无

返回值：

- bool 如果函数支持可变参数，返回true。否则返回false。

功能说明：

- reflect.TypeOf(interface{}).IsVariadic() 如果一个函数类型的最后一个输入参数是一个“...”参数，返回true。如果出现恐慌（panic），表示变量类型不是函数。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
		
	func main(){
		var a func(a1 int, a2 string, a3...interface{})
		var typeA reflect.Type = reflect.TypeOf(a)
		fmt.Println(typeA.IsVariadic())
		//>>true
		
		var b func(a1 int, a2 string, a3 interface{})
		var typeB reflect.Type = reflect.TypeOf(b)
		fmt.Println(typeB.IsVariadic())
		//>>false
	}
