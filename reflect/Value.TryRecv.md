# func (v Value) TryRecv() (x Value, ok bool)

参数列表

- 无

返回值：

- x Value 尝式从通道接收一个值，类型是 reflect.Value
- ok bool 信道能接收到值，返回true。信道被关闭，则返回false。

功能说明：

- reflect.ValueOf(interface{}).TryRecv(reflect.Value, bool) 尝式从通道 v 接收一个值，但不会阻塞。如果出现恐慌（panic），表示v的Kind不是Chan。如果x是零值，接收将不能完成而阻塞。如果x的值来自通道，布尔确定是true。如果是false，它收到是一个零值，因为通道被关闭。

代码实例：
	
	package main
	import (
		"fmt"
		"reflect"
	)
	
	func main(){
		var a chan int
		a = make(chan int,1)
		var value reflect.Value = reflect.ValueOf(&a)
		
		//判断指针是否指向内存地址
		if !value.CanSet() {
			value = value.Elem() //使指针指向内存地址
		}
		
		a<-1 // 值先发往信道
		var c, ok = value.TryRecv()
		fmt.Println(c.Int(), ok)
		//>>1 true
	}
