# type SelectDir int
参数列表

- 无

返回值：

- 无

功能说明：

- SelectDir 描述信道方向的选择case

		const (
			_	SelectDir = iota		// 初始化
		    SelectSend					// 发送方向 Chan <- Send
		    SelectRecv					// 接受方向 <-Chan:
		    SelectDefault				// 默认方向
		)


代码实例1：

	package main
	import (
		"fmt"
		"reflect"
	)
	  
	func main(){
		var selectCase = make([]reflect.SelectCase, 1)
		selectCase[0].Dir = reflect.SelectRecv //2
		fmt.Println(selectCase)
		//[{2 <chan int Value> <invalid Value>}]
		
		var selectCase1 = make([]reflect.SelectCase, 1)
		selectCase1[0].Dir = reflect.SelectSend //1
		fmt.Println(selectCase1)
		//[{1 <invalid Value> <chan int Value>}]
		
		var selectCase2 = make([]reflect.SelectCase, 1)
		selectCase2[0].Dir = reflect.SelectDefault //3
		fmt.Println(selectCase2)
		//[{3 <invalid Value> <invalid Value>}]
	}

