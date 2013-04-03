# ChanDir() ChanDir

参数列表

- 无

返回值：

- ChanDir 返回一个信道特定的 reflect.ChanDir 类型
	- reflect.ChanDir 有三个常量
		- reflect.SendDir 发送信道
		- reflect.RecvDir 接收信道
		- reflect.BothDir 双向信道

功能说明：

- reflect.TypeOf(interface{}).ChanDir() 返回信道类型的方向，如果出现恐慌（panic），那么它不是 Chan 类型。

代码实例：
	
	package main
	import (
	    "fmt"
	    "reflect"
	)
		
	func main(){
		var a chan<- int
		var typeA reflect.Type = reflect.TypeOf(a)
		fmt.Println(typeA.ChanDir(), typeA.ChanDir() == reflect.SendDir)
		//>>chan<- true
		
		var b <-chan int
		var typeB reflect.Type = reflect.TypeOf(b)
		fmt.Println(typeB.ChanDir(), typeB.ChanDir() == reflect.RecvDir)
		//>><-chan true
		
		var c chan int
		var typeC reflect.Type = reflect.TypeOf(c)
		fmt.Println(typeC.ChanDir(), typeC.ChanDir() == reflect.BothDir)
		//>>chan true
	}
