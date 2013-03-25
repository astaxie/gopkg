# func (v Value) TrySend(x Value) bool

参数列表

- x Value 往信道发送的值，类型是 reflect.Value

返回值：

- ok bool 如果值能往信道发送成功，返回true，否则返回false。

功能说明：

- reflect.ValueOf(interface{}).TrySend(reflect.Value) 尝试通道发送 x 给 v，但不会阻止。如果出现恐慌，表示v的Kind不是Chan。如果该值被发送，则返回true，否则返回false。在GO中，x的值必须是分配给该通道的元素类型。

代码实例：
	
	package main
	import (
		"fmt"
		"reflect"
	)
	
	func main(){
		var a chan int
		a = make(chan int,1)
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		var ok = value.TrySend(reflect.ValueOf(1)) //发送一个值往信道
		var c = <- a
		fmt.Println(c, ok)
		//>>1 true
	}
