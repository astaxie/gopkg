# func (v Value) Call(in []Value) []Value

参数列表

- in []Value 切片的值是 reflect.Value 类型

返回值：

- []Value 从函数内以切片类型返回的多个值
		
功能说明：

- reflect.ValueOf(interface{}).Call([]Value)  调用函数，in 切片装入函数的参数，并传到函数。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a = func(a1 int, a2 string, a3 float64) (int, string, float64){
			fmt.Println(a1, a2, a3)
			//>>123 456 123.456
			return a1, a2, a3
		}
		
		var value reflect.Value = reflect.ValueOf(a)
		var tt []reflect.Value
		tt = append(tt, reflect.ValueOf(123))
		tt = append(tt, reflect.ValueOf("456"))
		tt = append(tt, reflect.ValueOf(123.456))
		var tt1 []reflect.Value = value.Call(tt)
		fmt.Println(tt1[0], tt1[1], tt1[2])
		//>><int Value> 456 <float64 Value>
	}
