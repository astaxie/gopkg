# PkgPath() string

参数列表

- 无

返回值：

- string 包(package)的路径

功能说明：

- reflect.TypeOf(interface{}).PkgPath() 返回包(package)的路径

代码实例：
  
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		type A struct {}
		var a A
		var typeof reflect.Type = reflect.TypeOf(a)
		fmt.Println(typeof.PkgPath())
		//>>main

		var b int
		typeof = reflect.TypeOf(b)
		fmt.Println(typeof.PkgPath())
		//>>
	}
