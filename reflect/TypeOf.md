# func TypeOf(i interface{}) Type

参数列表

- i interface{} 是一个接口类型，可以传入任何类型变量

返回值：

- Type 返回反射类型 reflect.Type

功能说明：

- TypeOf 返回interface{}接口的值反射它的类型。 TypeOf(nil) 返回nil。

代码实例：
  
	package main
	import (
		"fmt"
		"reflect"
	)
	
	func main(){
		var a int
		var typeof reflect.Type = reflect.TypeOf(a)
		fmt.Println(typeof.String())
		//>>int
	}
