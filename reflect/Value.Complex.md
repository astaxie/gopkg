# func (v Value) Complex() complex128

参数列表

- 无

返回值：

- complex128  以最大复数的complex128类型返回值。

功能说明：

- reflect.ValueOf(interface{}).Complex() 返回 v 的值。如果出现恐慌，表示v的值不是Complex64 或Complex128 类型。

代码实例：
  
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a complex64
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		fmt.Println(value.Complex())
		//>>(0+0i)
		
		a = 1+5i
		
		fmt.Println(value.Complex())
		//>>(1+5i)
		
		fmt.Println(a)
		//>>(1+5i)
	}
