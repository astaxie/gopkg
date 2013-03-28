# func (v Value) OverflowInt(x int64) bool

参数列表

- x int64 传入参数必须是int64，以最大范围。

返回值：

- bool 值超出范围，返回true，否则返回false

功能说明：

- reflect.ValueOf(interface{}).OverflowInt(int64) 判断值是否溢出。如果出现恐慌，表示v的值不是Int，Int8，Int16，Int32，或Int64 类型。
	- 如果int16或Int32类型存储的值超出了该类型可容纳的范围，返回true，否则返回false。
	- 警告：如果int64类型存储的值超出了该类型可容纳的范围，返回错误，否则返回false。
	- int各类型存储范围：
		- Int8  值类型表示值介于 -128 到 +127 之间的有符号整数。
		- Int16 值类型表示值介于 -32,768 到 +32,767 之间的有符号整数。
		- Int32 值类型表示值介于 -2,147,483,648 到 +2,147,483,647 之间的有符号整数。
		- Int64 值类型表示值介于 -9,223,372,036,854,775,808 到 +9,223,372,036,854,775,807 之间的整数。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a int32 //注意：类型是32位
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		var i int64 = 2147483647
		if value.OverflowInt(i) {
			fmt.Println("溢出：超出范围")
		}else {
			fmt.Println("正常：范围之内")
		}
		//>>正常：范围之内
		
		i = 2147483648
		if value.OverflowInt(i) {
			fmt.Println("溢出：超出范围")
		}else {
			fmt.Println("正常：范围之内")
		}
		//>>溢出：超出范围
		
	}
