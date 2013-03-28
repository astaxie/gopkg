# func (v Value) Kind() Kind

参数列表

- 无

返回值：

- Kind 一种特定的 reflect.Kind 类型
		
功能说明：

- reflect.ValueOf(interface{}).Kind()  返回特定的一种类型 reflect.Kind，如果v是零值（IsValid返回false），返回无效的Kind。。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a string = "A"
		var value reflect.Value = reflect.ValueOf(a)
		fmt.Println(value.String(), value.Kind() == reflect.String)
		//>>A true
		
		var b int = 999
		var value1 reflect.Value = reflect.ValueOf(b)
		fmt.Println(value1.Int(), value1.Kind() == reflect.Int)
		//>>999 true
		
		var c float64 = 999.999
		var value2 reflect.Value = reflect.ValueOf(c)
		fmt.Println(value2.Float(), value2.Kind() == reflect.Float64)
		//>>999.999 true
	}
