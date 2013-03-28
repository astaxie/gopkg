# func (v Value) Bool()

参数列表

- 无

返回值：

- bool 布尔值

功能说明：

- reflect.ValueOf(interface{}).Bool() 返回 v 的布尔值。如果出现恐慌（panic），表示值的类型不是Bool。


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
		
		a = true
		
		fmt.Println(value.Bool())
		//>>true
	}
