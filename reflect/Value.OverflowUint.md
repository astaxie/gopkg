# func (v Value) OverflowUint(x uint64) bool

参数列表

- x uint64 传入参数必须是uint64，以最大范围。

返回值：

- bool 值超出范围返回true，否则返回false。

功能说明：

- reflect.ValueOf(interface{}).OverflowUint(uint64) 判断 Uint 值是否溢出。如果出现恐慌，表示v的值不是Uint，Uint8，Uint16，Uint32，或Uint64 类型。
	- 如果uint16或uint32类型存储的值超出了该类型可容纳的范围，返回true，否则返回false。
	- 警告：如果uint64类型存储的值超出了该类型可容纳的范围，返回错误，否则返回false。
		- uint各类型存储范围：
		- Uint8 值类型表示值介于 0 到 255 之间的无符号整数。也是有符号整数 128*2-1 。
		- Uint16 值类型表示值介于 0 到 65535 之间的无符号整数。也是有符号整数 32768*2-1 。
		- Uint32 值类型表示值介于 0 到 4294967295 之间的无符号整数。也是有符号整数 2147483648*2-1 。
		- Uint64 值类型表示值介于 0 到 18446744073709551615 之间的无符号整数。也是有符号整数 9223372036854775808*2-1 。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a uint32 //注意：类型是32位
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		var i uint64 = 4294967295 //等介于 2147483648*2-1
		if !value.OverflowUint(i) {
			fmt.Println("正常：范围之内")
		}else {
			fmt.Println("溢出：超出范围")
		}
		//>>正常：范围之内
		
		i = 4294967296 //等介于 2147483648*2-0
		if !value.OverflowUint(i) {
			fmt.Println("正常：范围之内")
		}else {
			fmt.Println("溢出：超出范围")
		}
		//>>溢出：超出范围
	}
