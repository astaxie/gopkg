# func (v Value) MethodByName(name string) Value

参数列表

- name string 传入方法的"字符串"名称

返回值：

- Value 返回结构绑定的方法 reflect.Value 类型

功能说明：

- reflect.ValueOf(interface{}).MethodByName(string)  指定一个“字符串”的方法名称返回方法的 reflect.Value 类型。如果出现恐慌(panic)，表示该类型不是struct。

代码实例：
  
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	type A struct {
	}
	func (p A) test(){}
	func (p A) test1(){}
	
	func main(){
		var a A
		var value reflect.Value = reflect.ValueOf(a)
		vf := value.MethodByName("test1")
		fmt.Println(vf.Kind(), vf.Kind() == reflect.Func)
		//>>func true
	}
