# func (v Value) Type() Type

参数列表

- 无

返回值：

- Type  返回一个特定的 reflect.Type 类型

功能说明：

- reflect.ValueOf(interface{}).Type() 将 Value 类型“转到” Type 类型。Type 类型也可以“转到”Value 类型上的，可惜只有 reflect.Method.Func 才有，reflect.StructField 类型中没有。

代码实例：
  
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a int
		var value reflect.Value = reflect.ValueOf(a)
		var typeof reflect.Type = value.Type() //reflect.Value 类型“转到” reflect.Type 类型
		fmt.Println(typeof)
	}
