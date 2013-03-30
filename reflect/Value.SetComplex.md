# func (v Value) SetComplex(x complex128)

参数列表

- x complex128 传入参数必须是complex128类型，以最大范围。

返回值：

- 无

功能说明：

- reflect.ValueOf(interface{}).SetComplex(complex128) 设置 v 的为 x。如果出现恐慌（panic），表示v的值不是Complex64 或Complex128 类型。
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
		var a complex64 //注意：类型是complex64
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		var i complex128 = -3.4028234663852884e+38 + 3.4028234663852884e+38i
		if !value.OverflowComplex(i) {
			value.SetComplex(i)
		}
		fmt.Println(value.Complex())
		//>>(-3.4028234663852886e+38+3.4028234663852886e+38i)
	}
