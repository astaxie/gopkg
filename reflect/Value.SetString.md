# func (v Value) SetString(x string)

参数列表

- x string 传入字符串

返回值：

- 无

功能说明：

- reflect.ValueOf(interface{}).SetString(string) 设置 v的值为x。如果出现恐慌（panic），表示v的Kind不是String字符串。

代码实例：
	
	package main
	import (
		"fmt"
		"reflect"
	)
	
	func main(){
		var a string = "A"
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		fmt.Println(value.String())
		//>>A
		
		value.SetString("B") //设置新字符
		
		fmt.Println(value.String())
		//>>B
	}
