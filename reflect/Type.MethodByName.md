# MethodByName(string) (Method, bool)

参数列表

- string 传入“字符串”的函数名称

返回值：

- .Method  返回函数的 reflect.Method 类型
- bool 布尔值，如果找到指定的函数，返回true，则返回false。

功能说明：

- reflect.TypeOf(interface{}).MethodByName(string) 使用“字符串”的函数名称返回函数的 reflect.Method 类型和一个布尔值。在struct结构中

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
		var m, ok = typeof.MethodByName("test1")
		fmt.Println(ok, m.Index, m.PkgPath, m.Name, m.Type, m.Func)
		//>>true 1 main test1 func(main.A) <func(main.A) Value>
	}
