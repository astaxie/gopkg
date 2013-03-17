# Name() string

参数列表

- 无

返回值：

- string 类型的名字

功能说明：

- reflect.TypeOf(interface{}).Name() 返回类型的名字

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a int
		var typeof reflect.Type = reflect.TypeOf(a)
		fmt.Println(typeof.Name())
		//>>int
	}
