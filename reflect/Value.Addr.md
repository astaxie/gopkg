# func (v Value) Addr() Value

参数列表

- 无

返回值：

- Value 返回一个指针值的地址
		
功能说明：

- reflect.ValueOf(interface{}).Addr() 返回一个指针值的地址，如果V 的 CanAddr() 返回false，它会出现恐慌（panic）。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a int
		var value reflect.Value = reflect.ValueOf(&a)
		var ve = value.Elem()
		if ve.CanAddr() {
			fmt.Println(ve.Addr(), ve.Kind(), ve.Int())
			//>><*int Value> int 0
		}
	}
