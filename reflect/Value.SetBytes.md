# func (v Value) SetBytes(x []byte)

参数列表

- x []byte 新的值（字节切片）

返回值：

- 无

功能说明：

- reflect.ValueOf(interface{}).SetBytes([]byte) 设置 v 的值。如果出现恐慌（panic)，表示v的值不是字节（Bytes）的切片（Slice）。

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
		
		value.SetBytes([]byte{'1','2','3','4','5'}) //设置新值
		
		fmt.Println(value.Bytes())
		//>>[49 50 51 52 53]
		
		fmt.Println(a)
		//[49 50 51 52 53]
	}
