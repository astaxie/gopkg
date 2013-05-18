# ConvertibleTo(u Type) bool
参数列表

- u Type 另一个变量的反映类型，类型是 reflect.Type

返回值：

- bool 返回true，如果一个类型的值可以转换为 u 类型。

功能说明：

- ConvertibleTo 返回true，如果一个类型的值可以转换为 u 类型。

代码实例：
  package main
	import (
		"fmt"
		"reflect"
	)
    
	func main(){
		var i int64
		var typeOf reflect.Type = reflect.TypeOf(i)
		
		var k string
		fmt.Println(typeOf.ConvertibleTo(reflect.TypeOf(k)))
		//>>true
		
		var k1 rune
		fmt.Println(typeOf.ConvertibleTo(reflect.TypeOf(k1)))
		//>>true
		
		var k2 uintptr
		fmt.Println(typeOf.ConvertibleTo(reflect.TypeOf(k2)))
		//>>true
		
		var k3 complex128
		fmt.Println(typeOf.ConvertibleTo(reflect.TypeOf(k3)))
		//>>false
	}
