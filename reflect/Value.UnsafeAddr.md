# func (v Value) UnsafeAddr() uintptr

参数列表

- 无

返回值：

- uintptr 返回指针（整数）
		
功能说明：

- reflect.ValueOf(interface{}).UnsafeAddr() 返回一个指针指向v的数据。如果v是不可寻址，会出现恐慌（panic）。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a int
		var value reflect.Value = reflect.ValueOf(&a)
		fmt.Println(value.Elem().UnsafeAddr(), uintptr(unsafe.Pointer(&a)))
		//>>282918960 282918960
	}
