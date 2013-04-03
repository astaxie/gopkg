# Size() uintptr

参数列表

- 无

返回值：

- uintptr 数据的字节数

功能说明：

- reflect.TypeOf(interface{}).Size() 返回值存储的字节数，是类似 unsafe.Sizeof。

代码实例：
  
	package main
	import (
	    "fmt"
	    "reflect"
		"unsafe"
	)
	
	func main(){
		var a int
		a = 1234
		var typeof reflect.Type = reflect.TypeOf(a)
		fmt.Println(typeof.Size(), unsafe.Sizeof(a)) //这两个功能是等介的
		//>>4 4
	}
