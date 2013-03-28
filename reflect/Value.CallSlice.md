# func (v Value) CallSlice(in []Value) []Value

参数列表

- in []Value 切片的值是 reflect.Value 类型

返回值：

- []Value 从函数内以切片类型返回的多个值，值是 reflect.Value 类型
		
功能说明：

- reflect.ValueOf(interface{}).CallSlice([]Value)  调用函数，in 切片装入参数，传到函数。用于可变参数函数。例如，if len(in) == 3，v.Call(in) 表示的go调用 test(in[0], in[1], in[2]...)。如果调用时出现恐慌（panic），表示类型的Kind不是Func，或者test不是可变参数。返回输出结果的Value。在GO中，每个输入参数必须是分配给函数的相应的输入参数的类型。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a = func(a1 int, a2 string, a3 ...float64) (int, string, []float64){
			fmt.Println(a1, a2, a3)
			//>>123 456 [1.1 2.2 3.3 4.4 5 5]
			return a1, a2, a3
		}
		
		var value reflect.Value = reflect.ValueOf(a)
		var tt []reflect.Value
		tt = append(tt, reflect.ValueOf(123))
		tt = append(tt, reflect.ValueOf("456"))
		tt = append(tt, reflect.ValueOf([]float64{1.1, 2.2, 3.3, 4.4, 5,5}))
		var tt1 []reflect.Value = value.CallSlice(tt)
		fmt.Println(tt1)
		//>>[<int Value> 456 <[]float64 Value>]
	}
