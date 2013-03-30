# func MakeChan(typ Type, buffer int) Value

参数列表

- typ Type 传入信道的 reflect.Type 类型
- buffer int 初始化信道的缓冲区大小

返回值：

- Value 输出已被初始的信道 reflect.Value 类型

功能说明：

- 指定类型初始化一个信道和设置缓冲区大小。

代码实例：
	
	package main
	import (
		"fmt"
		"reflect"
	)
	
	func main(){
		var a chan int
		
		var value reflect.Value = reflect.ValueOf(&a)
		
		value = reflect.Indirect(value) //使指针指向内存地址
		
		value = reflect.MakeChan(value.Type(), 99)
		
		fmt.Println(value.Kind(), value.Cap())
		//>>chan 99
	}
