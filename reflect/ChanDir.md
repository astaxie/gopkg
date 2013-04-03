# type ChanDir

参数列表

- 无

返回值：

- 无

功能说明：

- ChanDir 代表的信道类型方向。
  + reflect.ChanDir 有三个常量成员
		- reflect.RecvDir 接收数据
		- reflect.SendDir 发送数据
		- reflect.BothDir 双向信道

代码实例：

		package main
		import (
		    "fmt"
		    "reflect"
		)
		func main() {
			var a chan int
			func(c <-chan int) { // <-chan 代表是信道接收数据
				var chanDir = reflect.TypeOf(c).ChanDir()
				fmt.Println(chanDir == reflect.RecvDir) // reflect.RecvDir 常量代表信道接收数据
				//>>true
			}(a)
			func(c chan<- int) { // chan<- 代表是信道发送数据
				var chanDir = reflect.TypeOf(c).ChanDir()
				fmt.Println(chanDir == reflect.SendDir) // reflect.SendDir 常量代表信道发送数据
				//>>true
			}(a)
			func(c chan int) {	// chan 代表是双向信道
				var chanDir = reflect.TypeOf(c).ChanDir()
				fmt.Println(chanDir == reflect.BothDir) // reflect.BothDir 常量代表双向信道
				//>>true
			}(a)
		}
