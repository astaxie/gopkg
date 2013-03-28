# func (v Value) NumMethod() int

参数列表

- 无

返回值：

- int 返回结构绑定的方法数量

功能说明：

- reflect.ValueOf(interface{}).NumMethod() 绑定在结构中的方法数量。如果出现恐慌(panic)，表示该类型不是struct。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	type A struct {
	}
	func (p A) test(){}
	func (p A) test1(){}
	
	func main(){
		var a A
		var value reflect.Value = reflect.ValueOf(a)
		fmt.Println(value.NumMethod())
		//>>2
	}
