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

代码实例2：

    a := "abc"    // 字符串 "abc"
    b := (*uintptr)(unsafe.Pointer(&a))  // b 存储 a 的地址
	
    var c []byte // d 将 c 的结构用 reflect.SliceHeader 表示
    d := (*reflect.SliceHeader)((unsafe.Pointer(&c)))
    d.Cap = len(a)
    d.Len = len(a)
    d.Data = *b  // *b 存储字符串首元素地址
    fmt.Println(c)
    // [97 98 99]

代码实例2：

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

代码实例3：

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
		
		var t=&T{a:1,b:2,c:3}
		p := unsafe.Sizeof(*t)
		fmt.Println(int(p))
	   //12
		
	    sl := reflect.SliceHeader{
	        Data: uintptr(unsafe.Pointer(t)),
	        Len:  int(p),
	        Cap:  int(p),
	    }
		
	    b := *(*[]byte)(unsafe.Pointer(&sl))
	    fmt.Println(len(b))
	    //12
	    fmt.Println(t, b)
		//&{1 2 3} [1 0 0 0 2 0 0 0 3 0 0 0]
		
		//------------第一切片
		b[0] = 3
		
		//------------第二切片，下面的2就是表示二切片
		b[4] = 5
		b[5] = 1 //等介于 5 * 2 + 1
		
		//------------第三切片，下面的3就是表示三切片
		b[8] = 255 //最大255
		b[9] = 2 //等介于 255 * 3 + 2
		
		fmt.Println(t, b)
		//&{3 261 767} [3 0 0 0 5 1 0 0 255 2 0 0]
	}
