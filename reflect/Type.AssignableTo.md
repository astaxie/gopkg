# AssignableTo(u Type) bool

参数列表

- u reflect.Type 传入一个类型

返回值：

- bool 如果类型可以赋值，返回true，否则返回false。

功能说明：

- reflect.TypeOf(interface{}).AssignableTo(reflect.Type) 如果一个类型的值是可分配给u，返回true。

代码实例：
  
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a int
		var b int
		var c string
		var typeA reflect.Type = reflect.TypeOf(a)

		var typeB reflect.Type = reflect.TypeOf(b)
		fmt.Println(typeA.AssignableTo(typeB))
		//>>true

		var typeC reflect.Type = reflect.TypeOf(c)
		fmt.Println(typeA.AssignableTo(typeC))
		//>>false
	}
