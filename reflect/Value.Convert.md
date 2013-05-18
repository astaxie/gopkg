# func (v Value) Convert(t Type) Value
参数列表

- t Type 另一个类型的reflect.Type

返回值：

- Value 返回 reflect.Value 类型，更多方法可以查看reflect.Value 结构中绑定的方法

功能说明：

- Convert 返回被 v 转换为类型 t 。如果通常Go转换规则不容许值v转换为类型t，转换将恐慌(panic)。

代码实例：

	package main
	import (
		"fmt"
		"reflect"
	)
    
	func main(){
		var a int = 65 //A
		var b string = "string"
		var valueOf reflect.Value = reflect.ValueOf(a)
		var convert reflect.Value = valueOf.Convert(reflect.TypeOf(b))
		fmt.Println(convert.Kind(), convert.String())
		//string A
	}

