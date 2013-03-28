# func (v Value) CanSet() bool

参数列表

- 无

返回值：

- bool 如果v的值可以更改或写入新值，返回true，否则返回false。
  	
功能说明：

- reflect.ValueOf(interface{}).CanSet() 如果v的值可以更改或写入新值，返回true。


代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a int
		var value reflect.Value = reflect.ValueOf(&a)
		fmt.Println(value.CanSet())
		//>>false
	
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
	
		fmt.Println(value.CanSet())
		//>>true
	}
