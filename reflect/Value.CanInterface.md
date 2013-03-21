# func (v Value) CanInterface() bool

参数列表

- 无

返回值：

- bool 如果值支持接口方式读出，返回true，否则返回false。
		
功能说明：

- reflect.ValueOf(interface{}).CanInterface()  判断值是否可以做为 interface{} 类型读出

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	type A struct {
		A0 string
		A1 int
	}
	func (f A) test(){}
	func (f A) test1(){}
	
	func main(){
		var a A
		var value reflect.Value = reflect.ValueOf(a)
		for i:=0; i<value.NumField(); i++ {
			vf := value.Field(i)
			fmt.Println(vf.CanInterface())
			// 0 >>true
			// 1 >>true
		}
		for i:=0; i<value.NumMethod(); i++ {
			vf := value.Method(i)
			fmt.Println(vf.CanInterface())
			// 0 >>false
			// 1 >>false
		}
	}
