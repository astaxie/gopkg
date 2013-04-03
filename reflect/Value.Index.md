# func (v Value) Index(i int) Value

参数列表

- i int 切片的序号(整数)

返回值：

- Value 返回指定切片中的第i切片 reflect.Value 类型

功能说明：

- reflect.ValueOf(interface{}).Index(int)  返回Array或Slice类型的第i个切片。如果出现恐慌(panic)，表示该类型不是Array、Slice或i超出了范围。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	type A struct {
		A0 []string
	}
	func main(){
		var a A
		a.A0 = append(a.A0, []string{"a","b","c","d"}...) //等介于 "a","b","c","d"
		var value reflect.Value = reflect.ValueOf(a)
		vf := value.Field(0).Index(3)
		fmt.Println(vf.String(), vf.Kind() == reflect.String)
		//>>d true
	}
