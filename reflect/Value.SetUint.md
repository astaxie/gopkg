# func (v Value) SetUint(x uint64)

参数列表

- x uint64 传入参数必须是uint64，以最大范围。

返回值：

- 无

功能说明：

- reflect.ValueOf(interface{}).SetUint(uint64) 设置 v 的值为 x。如果出现恐慌，表示v的值不是Uint，Uint8，Uint16，Uint32，或Uint64 类型。
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
			value.SetUint(i) // 执行这里
		}
		fmt.Println(value.Uint(), a)
		//>>4294967295 4294967295
		
		i = 4294967296 //等介于 2147483648*2-0
		if !value.OverflowUint(i) {
			fmt.Println("溢出")
		}else{
			value.SetUint(0) // 执行这里
		}
		fmt.Println(value.Uint(), a)
		//>>0 0
		
		i = 4294967297 //等介于 2147483648*2-(-1)
		if value.OverflowUint(i) {
			value.SetUint(0) // 执行这里
		}
		fmt.Println(value.Uint(), a)
		//>>0 0
	}
