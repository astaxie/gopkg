# func (v Value) Elem() Value

参数列表

- 无

返回值：

- Value 返回一个直接指向内存地址的 reflect.Value 类型

功能说明：

- reflect.ValueOf(interface{}).Elem() 将指针地址指向内存地址，如果出现恐慌（panic），表示该类型不是 Ptr

代码实例：
	
	package main
	import (
		"fmt"
		"reflect"
	)
	
	func main(){
		var a int
		var value reflect.Value = reflect.ValueOf(&a)
		fmt.Println(value.Kind(), value.Elem().Kind())
		//>>ptr int
	}
