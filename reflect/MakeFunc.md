# func MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value
参数列表

- typ Type 一个未初化函数的方法值，类型是reflect.Type
- fn func(args []Value) (results []Value) 另一个函数，作用于对第一个函数参数操作。

返回值：

- Value 返回 reflect.Value 类型，更多方法可以查看reflect.Value 结构中绑定的方法

功能说明：

- MakeFunc 返回一个新的类型“函数”包含 fn 函数（绑定着fn函数）。
  - 注：这个函数的价值，还在摸索中...

代码实例：

	package main
	import (
		"fmt"
		"reflect"
	)
    
	func main(){
		var swap = func(in []reflect.Value) []reflect.Value {
			fmt.Println(in[0].Interface(), in[1].Interface())
			//1 2
			return []reflect.Value{in[1], in[0]}
		}
		var makeSwap = func(fptr interface{}) {
			var valueOf reflect.Value = reflect.Indirect( reflect.ValueOf(fptr))
			var v reflect.Value = reflect.MakeFunc(valueOf.Type(), swap)
			valueOf.Set(v)
			
			fmt.Println(valueOf)
			//<func(int, int) (int, int) Value>
			fmt.Println(v)
			//<func(int, int) (int, int) Value>
		}
		var intSwap func(int, int) (int, int)
		makeSwap(&intSwap)
		fmt.Println(intSwap(1, 2))
		//2 1
	}
