# Kind() Kind

参数列表

- 无

返回值：

- Kind 返回特定的 reflect.Kind 类型

功能说明：

- reflect.TypeOf(interface{}).Kind()  返回特定的 Kind 类型。

代码实例：

	package main
	import (
	    "fmt"
	    "reflect"
	)

	func main(){
		var a int
		var typeof reflect.Type = reflect.TypeOf(a)
		fmt.Println(typeof.Kind(), typeof.Kind() == reflect.Int)
		//>>int true
	}
