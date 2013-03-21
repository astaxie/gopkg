# func (v Value) Interface() (i interface{})

参数列表

- 无

返回值：

- i interface{} 返回一个接口类型的值
		
功能说明：

- reflect.ValueOf(interface{}).Interface()  以接口类型读出数据

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
		//字段的类型都是可以支持 Interface() 方法读出值
		for i:=0; i<value.NumField(); i++ {
			vf := value.Field(i)
			if vf.CanInterface() {
				fmt.Println(vf.Interface())
				// 0 >>0
				// 1 >>0
			}
		}
		
		//方法就不可使用 Interface() 方法读出值，因为方法没有值。
		for i:=0; i<value.NumMethod(); i++ {
			vf := value.Method(i)
			if !vf.CanInterface() {
				fmt.Println("vf.CanInterface() == false")
				// 0 >>vf.CanInterface() == false
				// 1 >>vf.CanInterface() == false
			}
		}
	}
