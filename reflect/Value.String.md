# func (v Value) String() string

参数列表

- 无

返回值：

- string 以字符串形式返回变量的值

功能说明：

- reflect.ValueOf(interface{}).String() 作为一个字符串返回 v 的值。String是一个特殊方法，由于Go的字符串的方法约定。不像其他的gettters，如果v的Kind不是String，它并不会恐慌（panic）。相反，它返回一个字符串的形式为“<T value>”，其中T是v的类型。

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
		
		a = "B"
		
		fmt.Println(value.String())
		//>>B
	}
