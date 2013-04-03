# Align() int

参数列表

- 无

返回值：

- int 变量在内存中对齐值

功能说明：

- reflect.TypeOf(interface{}).Align() 在分配在内存时的此类型的一个值（以字节为单位）的对齐。

代码实例：

	package main
	import (
	    "fmt"
	    "reflect"
	)

	func main(){
		var a int
		var typeof reflect.Type = reflect.TypeOf(a)
		fmt.Println(typeof.Align())
		//>>4
		var b string
		typeof = reflect.TypeOf(b)
		fmt.Println(typeof.Align())
		//>>4
		var c uint
		typeof = reflect.TypeOf(c)
		fmt.Println(typeof.Align())
		//>>4
	}
