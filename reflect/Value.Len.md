# func (v Value) Len() int

参数列表

- 无

返回值：

- int 长度

功能说明：

- reflect.ValueOf(interface{}).Len() 返回 v 的长度。如果出现恐慌，表示 v 的Kind不是Array，Chan，Map，Slice，或String。

代码实例：
	
	package main
	import (
		"fmt"
		"reflect"
	)
	
	func main(){
		var a []int
		a = make([]int, 9)
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		fmt.Println(value.Len())
		//>>9
	}
