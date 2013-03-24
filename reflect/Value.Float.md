# func (v Value) Float() float64

参数列表

- 无

返回值：

- float64  以最大浮点数的float64类型返回值。

功能说明：

- reflect.ValueOf(interface{}).Float() 返回 v 的值。如果出现恐慌，表示v的值不是Float64 或Float32 类型。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a float64
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		fmt.Println(value.Float())
		//>>0
		
		a = 1.2
		
		fmt.Println(value.Float())
		//>>1.2
		
		fmt.Println(a)
		//>>1.2
	}
