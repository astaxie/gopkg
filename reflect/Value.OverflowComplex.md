# func (v Value) OverflowComplex(x complex128) bool

参数列表

- x complex128 传入参数必须是complex128类型，以最大范围。

返回值：

- bool 如果判断值超出范围，返回true，否则返回false。

功能说明：

- reflect.ValueOf(interface{}).OverflowComplex(complex128) 判断 complex 是否溢出。如果出现恐慌（panic），表示v的值不是Complex64 或Complex128 类型。
	- 如果Complex64类型存储的值超出了该类型可容纳的范围，返回true，否则返回false。
	- 警告：如果Complex128类型存储的值超出了该类型可容纳的范围，返回错误，否则返回false。
		- Complex各类型存储范围：
		- Complex64 值类型表示值介于 -3.4028234663852884e+38 到 3.4028234663852884e+38  // 2**127 * (2**24 - 1) / 2**23
		- Complex128 值类型表示值介于 -1.7976931348623158e+308 到 1.7976931348623158e+308 // 2**1023 * (2**53 - 1) / 2**52

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a complex64
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		var i complex128 = 3.4028234663852884e+38 + 3.4028234663852884e+38i
		if value.OverflowComplex(i) {
			fmt.Println("溢出：超出范围")
		}else {
			fmt.Println("正常：范围之内")
		}
		//>>正常：范围之内
		
		i = 3.4028234663852885e+38 + 3.4028234663852884e+38i
		if value.OverflowComplex(i) {
			fmt.Println("溢出：超出范围")
		}else {
			fmt.Println("正常：范围之内")
		}
		//>>溢出：超出范围
	}
