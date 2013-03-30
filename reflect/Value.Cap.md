# func (v Value) Cap() int

参数列表

- 无

返回值：

- int 容量

功能说明：

- reflect.ValueOf(interface{}).Cap() 返回 v 的容量。如果出现恐慌，表示 v 的Kind不是Array，Chan，Slice。

代码实例：
	
	package main
	import (
		"fmt"
		"reflect"
	)
	
	func main(){
		var a []int
		a = make([]int, 9, 12)
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		fmt.Println(value.Cap())
		//>>12
	}
