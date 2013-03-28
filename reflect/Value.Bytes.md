# func (v Value) Bytes() []byte

参数列表

- 无

返回值：

- []byte 返回值的字节切片数据

功能说明：

- reflect.ValueOf(interface{}).Bytes() 返回 v 的值（字节切片）。如果出现恐慌（panic)，表示v的值不是一个字节的切片。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
	
	func main(){
		var a []byte
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		fmt.Println(value.Bytes())
		//>>[]
		
		a = []byte{1,2,3,4}
		
		fmt.Println(value.Bytes())
		//>>[1 2 3 4]
		
		fmt.Println(a)
		//[1 2 3 4]
	}
