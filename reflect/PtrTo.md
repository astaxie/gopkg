# func PtrTo(t Type) Type

参数列表

- t Type 是一个反射类型

返回值：

- Type 返回指针类型(reflect.Type)

功能说明：

- PtrTo 返回元素 t 的指针类型。例如，t 表示 Foo 类型，则 PtrTo(T) 代表 *Foo 类型。

代码实例：
	
	package main
	import (
		"fmt"
		"reflect"
	)

	func main(){
		var a int
		var pt reflect.Type = reflect.PtrTo(reflect.TypeOf(a))
		fmt.Println(pt.Kind(), pt.Elem().Kind()) // .Emel() 表示直接寻址
		// >>*int int
	}
