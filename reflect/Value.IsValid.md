# func (v Value) IsValid() bool

参数列表

- 无

返回值：

- bool 如果是零值返回true，否则返回false。
		
功能说明：

- reflect.ValueOf(interface{}).IsValid()  如果v代表一个值，返回true。如果v是零值，则返回false。零值不是变量初始化值。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a struct{}
		var value reflect.Value = reflect.ValueOf(a)
		fmt.Println(value.IsValid())
		//>>true
		
		var b interface{}
		var value1 reflect.Value = reflect.ValueOf(b)
		fmt.Println(value1.IsValid())
		//>>false
		
		var c int
		var value2 reflect.Value = reflect.ValueOf(c)
		fmt.Println(value2.IsValid())
		//>>true
		
		var d string
		var value3 reflect.Value = reflect.ValueOf(d)
		fmt.Println(value3.IsValid())
		//>>true
	}
