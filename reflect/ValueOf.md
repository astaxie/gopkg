# func ValueOf(i interface{}) Value

参数列表

- i interface{} 传入任何变量类型

返回值：

- Value  返回一个特定的 reflect.Value 类型

功能说明：

- reflect.ValueOf(interface{}) 反射返回一个存储在接口 i 的值。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a int
		var value reflect.Value = reflect.ValueOf(a)
		fmt.Println(value)
		//>><int Value>
	}
