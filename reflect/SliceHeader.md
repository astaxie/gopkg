# type SliceHeader struct
参数列表

- 无

返回值：

- 无

功能说明：

- 表示切片运行时。它不能被安全地使用，或者可移植。

  	type SliceHeader struct {
			Data	uintptr	// 指针
			Len		int		// 长度
			Cap		int		// 容量
		}


代码实例1：

	package main
	import (
		"fmt"
		"reflect"
		"unsafe"
	)
    
	func main(){
	    type T struct {
	        a int
	        b int
	        c int
	    }
		
	    t := &T{a: 1, b: 2, c: 3}
	    p := unsafe.Sizeof(*t)
	    println(int(p))
	    //12
	    sl := &reflect.SliceHeader{
	        Data: uintptr(unsafe.Pointer(t)),
	        Len:  int(p),
	        Cap:  int(p),
	    }
		
	    b := *(*[]byte)(unsafe.Pointer(sl))
	    println(len(b))
	    //12
	    fmt.Println(b)
		//[1 0 0 0 2 0 0 0 3 0 0 0]
		
	    b[0] = 7
	    b[4] = 5
	    b[8] = 8
		
	    fmt.Println(t)
		//&{7 5 8}
	}

