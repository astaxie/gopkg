# Bits() int

参数列表

- 无

返回值：

- int 返回类型的位大小

功能说明：

- reflect.TypeOf(interface{}).Bits() 返回类型的位大小，如果出现恐慌（panic），那么它的类型不是Int、Uint、Float或Complex。

代码实例：

	package main
	import (
	    "fmt"
	    "reflect"
	)

	func main(){
		var a int8
		var typeA reflect.Type = reflect.TypeOf(a)
		fmt.Println(typeA.Bits())
		//>>8

		var b int64
		var typeB reflect.Type = reflect.TypeOf(b)
		fmt.Println(typeB.Bits())
		//>>64
	}
