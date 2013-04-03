# func (v Value) OverflowFloat(x float64) bool

参数列表

- float64 传入参数必须是Float64，以最大范围。

返回值：

- bool 如果float类型的值范出范围，返回true，否则返回false。

功能说明：

- reflect.ValueOf(interface{}).OverflowFloat(float64) 判断float 溢出。如果出现恐慌，表示v的值不是Float64 或Float32 类型。
	- 如果float32类型存储的值超出了该类型可容纳的范围，返回true，否则返回false。
	- 警告：如果float64类型存储的值超出了该类型可容纳的范围，返回错误，否则返回false。
		- float各类型存储范围：
		- Float32 值类型表示值介于 -3.4028234663852884e+38 到 3.4028234663852884e+38  // 2**127 * (2**24 - 1) / 2**23
		- Float64 值类型表示值介于 -1.7976931348623158e+308 到 1.7976931348623158e+308 // 2**1023 * (2**53 - 1) / 2**52

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a float32 //注意：类型是float32
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		var i float64 = 3.4028234663852884e+38
		if value.OverflowFloat(i) {
			fmt.Println("溢出：超出范围")
		}else {
			fmt.Println("正常：范围之内")
		}
		//>>正常：范围之内
		
		i = 1.7976931348623158e+308
		if value.OverflowFloat(i) {
			fmt.Println("溢出：超出范围")
		}else {
			fmt.Println("正常：范围之内")
		}
		//>>溢出：超出范围
	}
