# func NewAt(typ Type, p unsafe.Pointer) Value

参数列表

- typ Type 传入新值的 reflect.Type 类型
- p unsafe.Pointer 另一个新值的指针

返回值：

- Value 输出已被初始新值的指针，类型是 reflect.Value

功能说明：

- 返回一个值，该值表示指定类型的指针的值，使用p作为该指针。

代码实例：
	
	package main
	import (
		"fmt"
		"reflect"
		"unsafe"
	)
	
	func main(){
		var a []int
		var b []int
		var value reflect.Value = reflect.ValueOf(&a)
		var value1 reflect.Value = reflect.ValueOf(&b)
		
		value = reflect.Indirect(value) //使指针指向内存地址
		value1 = reflect.Indirect(value1) //使指针指向内存地址
		
		value1 = reflect.NewAt(value.Type(), unsafe.Pointer(reflect.ValueOf(&b).Pointer())).Elem()
		b = append(b, 1)
		a = append(a, 2)
		fmt.Println(value.Pointer(), value1.Pointer())
		//>>282918952 282918960
		
		value.Set(value1)
		
		fmt.Println(value1.Kind(), value.Pointer(), value1.Pointer(), value.Interface(), value1.Interface(), a, b)
		//>>slice 282918960 282918960 [1] [1] [1] [1]
	}
