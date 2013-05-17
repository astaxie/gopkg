# func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)
参数列表

- cases []SelectCase 使用此参数之前要设置SelectCase，告诉Select目的。

返回值：

- chosen int 返回参数[]SelectCase 切片的序号
- recv Value 返回信道的值，类型是reflect.Value
- recvOk bool 返回false，如果信道被关闭。

功能说明：

- Select 等介于 select，拥有一样的随机的选择。

代码实例1：

    package main
    import (
        "fmt"
        "reflect"
    )
    
    func main(){
  	var c = make( chan int, 1)
		go func(c chan int){
			for i:=0;i<10;i++{
				c<-i
			}
			close(c)
		}(c)
		
		L:for {
			select {
			case val, ok:=<-c:
				if ok {
					fmt.Println("接收: ", val)
				}else{
					fmt.Println("信道关闭了")
					break L
				}
			}
		}
		//>>接收:  0
		//>>接收:  1
		//>>接收:  2
		//>>接收:  3
		//>>接收:  4
		//>>接收:  5
		//>>接收:  6
		//>>接收:  7
		//>>接收:  8
		//>>接收:  9
		//>>信道关闭了
    }

代码实例2：

		var chs = make(chan int)
		
		var worker = func(c chan int){
			for i:=0;i<5;i++{
				c<-i
			}
			close(c)
		}
		
		go worker(chs)
		
		var selectCase = make([]reflect.SelectCase, 1)
		selectCase[0].Dir = reflect.SelectRecv //设置信道是接收
		selectCase[0].Chan = reflect.ValueOf(chs)
		
		numDone := 0
		for numDone < 1 {
			chosen, recv, recvOk := reflect.Select(selectCase)
			if recvOk {
				fmt.Println(chosen, recv.Int(), recvOk)
				//0 0 true
				//0 1 true
				//0 2 true
			}else{
				numDone++
			}
		}

代码实例3：

		var n int = 1
		var chs = make([]chan int, n)
		
		var worker = func(n int, c chan int){
			for i:=0;i<n;i++{
				c<-i
			}
			close(c)
		}
		
		for i:=0;i<n;i++{
			chs[i]=make(chan int)
			go worker(3+i, chs[i])
		}
		
		var selectCase = make([]reflect.SelectCase, n)
		for i:=0;i<n;i++{
			selectCase[i].Dir = reflect.SelectRecv //设置信道是接收
			selectCase[i].Chan = reflect.ValueOf(chs[i])
		}
		
		numDone := 0
		for numDone < n {
			chosen, recv, recvOk := reflect.Select(selectCase)
			if recvOk {
				fmt.Println(chosen, recv.Int(), recvOk)
				//0 0 true
				//0 1 true
				//0 2 true
			}else{
				numDone++
			}
		}
