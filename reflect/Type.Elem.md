# Elem() Type

参数列表

- 无

返回值：

- Type reflect.Type 返回一个直接指向内存地址的 reflect.Type 类型

功能说明：

- reflect.TypeOf(interface{}).Elem() 返回一个类型的元素类型。如果出现恐慌（panic），表示变量类型不是Array，Chan，Map，Ptr，或Slice。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a int
		var typeA reflect.Type = reflect.TypeOf(&a)
		fmt.Println(typeA, typeA.Kind()) // 这里是指针，间接指向内存地址
		//>>*int ptr
		
		var b int
		var typeB reflect.Type = reflect.TypeOf(&b)
		fmt.Println(typeB.Elem(), typeB.Elem().Kind()) // 加上.Elem() 就可以直接指向内存地址。Elem()很有用，很多时候都用的到它。
		//>>int int
		
		var c map[int]string
		var typeC reflect.Type = reflect.TypeOf(c)
		fmt.Println(typeC.Elem(), typeC.Elem().Kind())
		//>>string string
		
		var d map[int]string
		var typeD reflect.Type = reflect.TypeOf(&d)
		fmt.Println(typeD.Elem(), typeD.Elem().Elem(), typeD.Elem().Kind())
		//>>map[int]string string map
	}
