# func (v Value) Uint() uint64

参数列表

- 无

返回值：

- uint64 传入参数必须是uint64，以最大范围。

功能说明：

- reflect.ValueOf(interface{}).Uint() 返回 v 的值。如果出现恐慌，表示v的值不是Uint，Uint8，Uint16，Uint32，或Uint64 类型。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a uint
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		fmt.Println(value.Uint())
		//>>0
		
		a = 123
		
		fmt.Println(value.Uint())
		//>>123
		
		fmt.Println(a)
		//>>123
		
	}
