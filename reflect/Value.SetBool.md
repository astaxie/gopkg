# func (v Value) SetBool(x bool)

参数列表

- bool 布尔值

返回值：

- 无

功能说明：

- reflect.ValueOf(interface{}).SetBool(bool) 设置 v 的值。如果CanSet()返回False，表示不支持SetBool设置。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a bool // 默认值为 false
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		fmt.Println(value.Bool())
		//>>false
		
		value.SetBool(true) //设置新值
		
		fmt.Println(value.Bool())
		//>>true
		
		fmt.Println(a)
		//true
	}
