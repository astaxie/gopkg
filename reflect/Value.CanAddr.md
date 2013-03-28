# func (v Value) CanAddr() bool

参数列表

- 无

返回值：

- bool 如果该值的地址可以得到寻址， 返回true。否则返回false。
  	
功能说明：

- reflect.ValueOf(interface{}).CanAddr()  如果该值的地址可以得到寻址，返回true。这样的值被称为寻址。如果元素它是一个切片，值是可寻址的。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a  int
		var value reflect.Value = reflect.ValueOf(&a)
		fmt.Println(value.CanAddr(), value.Elem().CanAddr())
		//>>false true
	}
