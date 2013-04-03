# func (v Value) Recv() (x Value, ok bool)

参数列表

- 无 

返回值：

- x Value 信道返回的值
- ok bool 如果信道能返回值，返回true。如果信道被关闭，返回false。

功能说明：

- reflect.ValueOf(interface{}).Recv(reflect.Value, bool) 从通道 v 接收与返回一个值，如果出现恐慌，表示Kind不是Chan。

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
