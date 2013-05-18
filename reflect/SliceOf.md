# func SliceOf(t Type) Type
参数列表

- t Type 切片的类型，如[]int，传入int的反射类型。

返回值：

- Type 返回 reflect.Type 类型，更多方法可以查看reflect.Type 接口

功能说明：

- SliceOf 返回t元素类型的一个切片 。例如，如果t表示整型（int），SliceOf(t)就是表示[]int。

代码实例：

	package main
	import (
		"fmt"
		"reflect"
	)
    
	func main(){
		var i int
		var sliceOf reflect.Type = reflect.SliceOf(reflect.TypeOf(i))
		fmt.Println(sliceOf.Kind(), sliceOf.String())
		//slice []int
	}
