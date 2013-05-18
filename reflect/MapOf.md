# func MapOf(key, elem Type) Type
参数列表

- key Type 表示着Map[int]string 中的 int类型
- elem Type 表示着Map[int]string 中的 string类型

返回值：

- Type 返回 reflect.Type 类型，更多方法可以查看reflect.Type 接口

功能说明：

- MapOf 返回Map类型与键和元素的类型。例如，如果 k表示 int和 e表示 string，MapOf(k, e) 就是这样表示 Map[int]string。
- 如果 Key类型非是一个有效Map的 Key类型（也就是说，如果它没有实现Go的 == 操作符），MapOf恐慌（panic）。

代码实例：

	package main
	import (
		"fmt"
		"reflect"
	)
    
	func main(){
		var key int
		var val string
		var mapOf reflect.Type = reflect.MapOf(reflect.TypeOf(key), reflect.TypeOf(val))
		fmt.Println(mapOf.Kind(), mapOf.String())
		//map map[int]string
	}
