# func New(typ Type) Value

参数列表

- typ Type 传入新值的 reflect.Type 类型

返回值：

- Value 输出已被初始新值的 reflect.Value 类型

功能说明：

- 指定类型初始化一个新的零值，返回一个指针指向被初始化的零值Value。也就是说，返回值的类型是PtrTo(t)。

代码实例：
	
	package main
	import (
		"fmt"
		"reflect"
	)
	
	func main(){
		var a []int
		
		var value reflect.Value = reflect.ValueOf(&a)
		
		value = reflect.Indirect(value) //使指针指向内存地址
		
		value = reflect.New(value.Type()) //初始化后返回指针
		fmt.Println(value.Kind())
		//>>ptr
		
		value = reflect.Indirect(value) //使指针指向内存地址
		fmt.Println(value.Kind())
		//>>slice
	}
