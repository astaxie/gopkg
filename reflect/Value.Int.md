# func (v Value) Int() int64

参数列表

- 无

返回值：

- int64 以最大整数的int64类型返回值。

功能说明：

- reflect.ValueOf(interface{}).Int() 以最大整数的int64类型返回 v 的值。如果出现恐慌（panic），表示v的值不是Int，Int8，Int16，Int32，或Int64 类型。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a int
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		fmt.Println(value.Int())
		//>>0
		
		a = 123
		
		fmt.Println(value.Int())
		//>>123
		
		fmt.Println(a)
		//123
	}
