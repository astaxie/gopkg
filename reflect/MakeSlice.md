# func MakeSlice(typ Type, len, cap int) Valu

参数列表

- typ Type 传入Slice的 reflect.Type 类型
- len int 传入需要初始化切片的长度
- cap int 传入需要被始化切片的容量

返回值：

- Value 输出已被初始Map的 reflect.Value 类型

功能说明：

- 指定切片类型创建一个新的初始化切片长度和容量。

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
		
		value = reflect.MakeSlice(value.Type(), 88, 99)
		
		fmt.Println(value.Kind(), value.Len(), value.Cap())
		//>>slice 88 99
	}
