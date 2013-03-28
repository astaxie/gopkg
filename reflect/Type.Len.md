# Len() int

参数列表

- 无

返回值：

- int 返回数组Array的长度

功能说明：

- reflect.TypeOf(interface{}).Len() 返回数组Array的长度。如果出现恐慌（panic），表示变量类型不是Array。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a [33]int
		var typeA reflect.Type = reflect.TypeOf(a)
		fmt.Println(typeA.Len())
		//>>33
	}
