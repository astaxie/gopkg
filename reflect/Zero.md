# func Zero(typ Type) Value

参数列表

- typ Type 传入变量，类型是 reflect.Type 

返回值：

- Value 返回类型的零值，类型是 reflect.Value

功能说明：

- 返回一个值，该值表示为指定的类型的零值。

代码实例：
	
	package main
	import (
		"fmt"
		"reflect"
	)
	
	func main(){
		var a int
		a = 123456
		var value reflect.Value = reflect.ValueOf(&a)
		value = reflect.Indirect(value)
		value = reflect.Zero(value.Type()) //返回a的默认值
		
		fmt.Println(value.Kind(), value.Int(), a)
		//>>int 0 123456
	}
