# func (v Value) MapKeys() []Value

参数列表

- 无

返回值：

- []Value 返回一个切片包含的所有键Key

功能说明：

- reflect.ValueOf(interface{}).MapKeys() 返回一个切片包含的所有键Key，无顺序。如果出现恐慌（panic），表示 v 的Kind不是Map。如果v是一个零的Map，它返回一个空切片。

代码实例：
	
	package main
	import (
		"fmt"
		"reflect"
	)
	
	func main(){
		var a = map[string]int{
			"a":1,
			"b":2,
			"c":3,
			"d":4,
			"e":5,
		}
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		fmt.Println(value.MapKeys()) //无顺序的返回Key
		//>>[c e a d b]
	}
