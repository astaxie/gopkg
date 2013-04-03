# func (v Value) Method(i int) Value

参数列表

- i int 传入方法的索引号

返回值：

- Value 返回结构绑定的方法 reflect.Value 类型

功能说明：

- reflect.ValueOf(interface{}).Method(int) 指定方法索引号返回结构绑定的方法 reflect.Value 类型。如果出现恐慌(panic)，表示该类型不是struct。

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
		for i:=0; i<value.NumMethod(); i++ {
			vf := value.Method(i)
			fmt.Println(vf.Kind(), vf.Kind() == reflect.Func)
			// 0 >>func true
			// 1 >>func true
		}
	}
