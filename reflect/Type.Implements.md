# Implements(u Type) bool

参数列表

- u reflect.Type 一个结构字段的 Type

返回值：

- bool 如果结构能实现 u 的接口，返回true，否则返回false。

功能说明：

- reflect.TypeOf(interface{}).Implements(reflect.Type)  如果该结构类型实现 U 接口，返回true。

代码实例：

  package main
	import (
	    "fmt"
	    "reflect"
	)

	type A struct {
		C
	}
	type B struct {
		C
		D
		E
	}
	type C interface {
		test()
	}
	type D interface {
		test1()
	}
	type E interface {}

	func main(){
		var a A
		var b B
		var typeof reflect.Type = reflect.TypeOf(a)

		//判断a结构是否存在b结构中的C接口
		var booL = typeof.Implements( reflect.TypeOf(b).Field(0).Type )
		fmt.Println(booL)
		//>>true

		//判断a结构是否存在b结构中的D接口
		booL = typeof.Implements( reflect.TypeOf(b).Field(1).Type )
		fmt.Println(booL)
		//>>false

		//b结构中的E因为是空接口，所以也返回true。
		booL = typeof.Implements( reflect.TypeOf(b).Field(2).Type )
		fmt.Println(booL)
		//>>true
	}
