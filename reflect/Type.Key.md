# Key() Type

参数列表

- 无

返回值：

- Type reflect.Type 返回 map 类型 key 的类型。

功能说明：

- reflect.TypeOf(interface{}).Key() 返回 map 类型 key 的类型。如果出现恐慌（panic），表示变量类型不是Map。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a map[int]string
		var typeA reflect.Type = reflect.TypeOf(a)
		fmt.Println(typeA.Key(), typeA.Elem().Kind())
		//>>int string
	}
