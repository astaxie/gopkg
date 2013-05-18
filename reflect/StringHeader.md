# type StringHeader struct
参数列表

- 无

返回值：

- 无

功能说明：

- 表示切片运行时。它不能被安全地使用，或者可移植。

		type StringHeader struct {
			Data uintptr	// 指针
			Len  int		// 长度
		}

代码实例：
	package main
	import (
		"fmt"
		"reflect"
		"unsafe"
	)
    
	func main(){
	    var b = []byte{97, 98, 99}
	    fmt.Println(unsafe.Sizeof(b), len(b))
		//12 3
		
	    var sh = (*reflect.SliceHeader)(unsafe.Pointer(&b))
		fmt.Println(sh.Data, sh.Len, sh.Cap)
		//820188969 3 3
		fmt.Println((*[]byte)(unsafe.Pointer(&sh.Data)))
		//&[97 98 99]
		
		var s string
		var sh1 =  (*reflect.StringHeader)(unsafe.Pointer(&s))
		//sh1.Data = *(*uintptr)(unsafe.Pointer(&b))
		sh1.Data = sh.Data
		sh1.Len = sh.Len
		fmt.Println(s)
		//abc
	}
