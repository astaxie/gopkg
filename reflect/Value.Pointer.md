# func (v Value) Pointer() uintptr

参数列表

- 无

返回值：

- uintptr 返回指针值（整数）

功能说明：

- reflect.ValueOf(interface{}).Pointer() 返回 v 的指针值（uintptr）

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
			"unsafe"
	)
	
	func main(){
		var a complex64
		var value reflect.Value = reflect.ValueOf(&a)
		fmt.Println(value.Pointer(), uintptr(unsafe.Pointer(&a)))
	}
