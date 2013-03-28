# func (v Value) Close()

参数列表

- 无

返回值：

- 无

功能说明：

- reflect.ValueOf(interface{}).Close() 关闭 v 的信道。如果出现恐慌，表示v的Kind不是Chan。

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
		
		var gotest = func(c chan int, v *reflect.Value) {
			c<-1
			v.Close() //信道关闭
		}
		go gotest(a,&value)
		 var v ,ok = <-a
		fmt.Println(v, ok)
		//>>1 true
	}
