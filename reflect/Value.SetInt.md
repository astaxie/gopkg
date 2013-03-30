# func (v Value) SetInt(x int64)

参数列表

-  x int64 传入参数必须是int64，以最大范围。

返回值：

- 无

功能说明：

- reflect.ValueOf(interface{}).SetInt(x int64) 设置 v 的值为 x。如果出现恐慌，表示v的值不是Int，Int8，Int16，Int32，或Int64 类型。
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
		if !value.OverflowInt(i) {
			value.SetInt(i) //执行这里
		}
		fmt.Println(value.Int())
		//>>2147483647
		
		i = 2147483648
		if !value.OverflowInt(i) {
			value.SetInt(i)
		}else{
			value.SetInt(0) //执行这里
		}
		fmt.Println(value.Int())
		//>>0
		
		i = 2147483650
		if value.OverflowInt(i) {//条件不带！
			value.SetInt(i) // 写入超出范围的整，看看会是什么结果！！！
		}
		fmt.Println(value.Int())
		//>>-2147483646
		
	}
