# String() string

参数列表

- 无

返回值：

- string 字符串表示的类型

功能说明：

- reflect.TypeOf(interface{}).String() 返回一个字符串表示的类型

代码实例：
  
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a int
		var typeof reflect.Type = reflect.TypeOf(a)
		fmt.Println(typeof.String())
		//>>int
	}
