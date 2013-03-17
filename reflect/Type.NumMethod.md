# NumMethod() int

参数列表

- 无

返回值：

- int 返回函数的总数量

功能说明：

- reflect.TypeOf(x).NumMethod() 返回struct结构绑定的方法集。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	type A struct {
		A0 int
	}
	func (f A) test(){}
	func (f A) test1(){}
	func (f A) test2(){}
	
	func main(){
		var a A
		var typeof reflect.Type = reflect.TypeOf(a)
		fmt.Println(typeof.NumMethod())
		//>>3
	}
