# func (v Value) InterfaceData() [2]uintptr

参数列表

- 无

返回值：

- [2]uintptr 返回一对作为uintptr的接口值
		
功能说明：

- reflect.ValueOf(interface{}).InterfaceData()  返回一对作为uintptr的接口值，如果出现恐慌，表示该类型没有接口。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	type A struct {
		A0 float64
		A1 B
	}
	type B interface {}
	
	func main(){
		var a A
		var value reflect.Value = reflect.ValueOf(a)
		fmt.Println(value.Field(1).InterfaceData())
		//>>[0 0]
	}
