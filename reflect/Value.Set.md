# func (v Value) Set(x Value)

参数列表

- x Value 新的值，类型为 reflect.Value

返回值：

- 无

功能说明：

- reflect.ValueOf(interface{}).Set(reflect.Value) 更改 v 的值为 x，支持所有类型的值更改。 如果CanSet()返回False，表示v的值不支持Set设置。


代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	type A struct {
		A0 float64
		A1 int
		A2 string
	}
	
	func main(){
		var a A
		var value reflect.Value = reflect.ValueOf(&a)
		
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		//更改a结构字段的值
		value.Field(0).Set(reflect.ValueOf(123.123))
		value.Field(1).Set(reflect.ValueOf(456))
		value.Field(2).Set(reflect.ValueOf("789"))
		
		fmt.Println(a)
		//>>{123.123 456 789}
	}
