# func Indirect(v Value) Value

参数列表

- v Value 传入是一个 reflect.Value 类型的变量

返回值：

- Value 输出v的指针

功能说明：

- 返回值v的指针。如果v是一个零指针，间接返回零值。如果v不是一个指针，间接的返回v。

代码实例：
	
	package main
	import (
		"fmt"
		"reflect"
	)
	
	func main(){
		var a []int
		
		//例子1
		var value reflect.Value = reflect.ValueOf(&a)
		fmt.Println(value.Kind())
		//>>ptr
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		fmt.Println(value.Kind())
		//>>slice
		
		//例子2
		var value1 reflect.Value = reflect.ValueOf(&a)
		fmt.Println(value1.Kind())
		//>>ptr
		
		value1 = reflect.Indirect(value1) //功能等介于上面例子1
		
		fmt.Println(value1.Kind())
		//>>slice
	}
