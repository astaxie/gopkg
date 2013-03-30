# func (v Value) MapIndex(key Value) Value

参数列表

- key Value 传入Map的键Key，类型是 reflect.Value

返回值：

- Value 返回Map的值，类型是 reflect.Value

功能说明：

- reflect.ValueOf(interface{}).MapIndex(reflect.Value) 指定键Key返回在v上的Map值。如果出现恐慌（panic），表示v的Kind不是Map。如果key在Map上没有找到，或v的Map是一个零Map，它返回零值。在Go中，该键的值必须是分配给map的键类型。

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
		
		var(
			mv reflect.Value
			keys []reflect.Value = value.MapKeys() //返回无顺序的Keys
		)
		for i := 0; i<len(keys); i++ {
			mv = value.MapIndex(keys[i]) //使用Key名称指定返回值
			fmt.Println(keys[i].String(), mv.Int())
			// 0 >>c 3
			// 1 >>d 4
			// 2 >>e 5
			// 3 >>a 1
			// 4 >>b 2
		}
	}
