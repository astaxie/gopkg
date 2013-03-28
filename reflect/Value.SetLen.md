# func (v Value) SetLen(n int)

参数列表

- n int 限制Slice的长度

返回值：

- 无

功能说明：

- reflect.ValueOf(interface{}).SetLen(int) 设置 v 的长度为 n。如果出现恐慌，表示 v 的Kind不是Slice或n为负数或超过Slice的容量。

代码实例：
	
	package main
	import (
		"fmt"
		"reflect"
	)
	
	func main(){
		var a []int
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		value.Set(reflect.ValueOf([]int{1,2,3,4,5,6,7,8,9,0}))
		
		fmt.Println(value.Interface())
		//>>[1 2 3 4 5 6 7 8 9 0]
		
		value.SetLen(5) //限制Slice长度
		
		fmt.Println(value.Interface())
		//>>[1 2 3 4 5]
	}
