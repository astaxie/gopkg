# func (v Value) Send(x Value)

参数列表

- x Value 发送数据往信道

返回值：

- 无

功能说明：

- reflect.ValueOf(interface{}).Send(reflect.Value) 发送x往通道上。 如果出现恐慌（panic），表示v的Kind不是Chan。在Go中，x的值必须是分配给该通道的元素类型。

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
		var gotest = func(c reflect.Value){
			c.Send(reflect.ValueOf(3888)) //3>发送信息
		}
		go gotest(value) //1>异步
		
		var c, ok = value.Recv() //2>阻塞等待
		
		fmt.Println(c.Int(), ok) //4>读出数据
		//>>3888 true
	}
