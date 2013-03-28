# func (v Value) IsNil() bool

参数列表

- 无

返回值：

- bool 如果是空值返回true，否则返回false。
		
功能说明：

- reflect.ValueOf(interface{}).IsNil()  判断值是否是 nil，限制支持 Chan，Func，Interface，Map，Ptr，或Slice。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a chan int
		var value reflect.Value = reflect.ValueOf(&a)
		fmt.Println(value.IsNil())
		//>>false
		
		var b []uintptr
		var value1 reflect.Value = reflect.ValueOf(b)
		fmt.Println(value1.IsNil())
		//>>true
		
		b = append(b, 123) //赋值 123
		value1 = reflect.ValueOf(b)
		fmt.Println(value1.IsNil()) // 赋值后，不再是空值
		//>>false
		
		var c func()
		var value2 reflect.Value = reflect.ValueOf(c)
		fmt.Println(value2.IsNil())
		//>>true
		
		var d map[int]string
		var value3 reflect.Value = reflect.ValueOf(d)
		fmt.Println(value3.IsNil())
		//>>true

	}
