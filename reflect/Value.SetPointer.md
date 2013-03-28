# func (v Value) SetPointer(x unsafe.Pointer)

参数列表

- x unsafe.Pointer 传入一个unsafe.Pointer的值

返回值：

- 无

功能说明：

- reflect.ValueOf(interface{}).SetPointer(unsafe.Pointer) 设置v的unsafe.Pointer值为 x 。如果出现恐慌，表示 v 不是UnsafePointer。

代码实例：
  
	package main
	import (
		"fmt"
		"reflect"
		"unsafe"
	)
	
	func main(){
		var a int = 456
		var ap unsafe.Pointer = unsafe.Pointer(&a)
		var value reflect.Value = reflect.ValueOf(&ap).Elem()
		
		var b int = 123
		var bp unsafe.Pointer = unsafe.Pointer(&b)
		
		fmt.Println(value.Pointer(), value.Kind(), uintptr(ap), uintptr(bp), a, b)
		//>>282918960 unsafe.Pointer 282918960 282918944 456 123
		
		value.SetPointer(bp)
		
		fmt.Println(value.Pointer(), value.Kind(), uintptr(ap), uintptr(bp), a, b)
		//>>282918944 unsafe.Pointer 282918944 282918944 456 123
	}
