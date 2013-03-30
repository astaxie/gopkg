# func (v Value) SetMapIndex(key, val Value)

参数列表

- key Value 传入对应的键名称，类型是 reflect.Value
- val Value 传入一个键的新值，类型是 reflect.Value

返回值：

- 无

功能说明：

- reflect.ValueOf(interface{}).SetMapIndex(reflect.Value) 设置 v 的Map键值。如果 v 的Kind不是Map，他会恐慌（panic）。如果值是零值（nil），SetMapIndex从Map上删除该键。在Go中，Key必须是分配给map的Key类型，和value必须是分配到Map的value类型中。

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
		
		value.SetMapIndex(reflect.ValueOf("c"), reflect.ValueOf(123456)) //修改键值
		
		value.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(nil)) //如果值是nil，则删除该键
		
		value.SetMapIndex(reflect.ValueOf("f"), reflect.ValueOf(789)) //增加一个键
		
		fmt.Println(value.Interface())
		//>>map[f:789 b:2 e:5 d:4 c:123456]
	}
