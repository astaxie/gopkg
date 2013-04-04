#WaitGroup

WaitGroup的结构如下：

	type WaitGroup struct {
		m		Mutex
		counter	int32
		waiters	int32
		sema	*uint32
	}
	
##成员

-	m：防止对内部成员进行竞争操作。
-	counter：需要等待的过程数。
-	waiters：调用`Wait()`方法的次数，当counter==0时，全部唤醒。
-	sema：用于waiter的信号量。

## func (wg *WaitGroup) Add(delta int)

参数：

-	int类型。

返回值：

-	无

功能：

-	添加delta个元素进入WaitGroup。个人觉得这里的接口设计有问题，如果Add(2)，那就要掉用两次Done()。而且会有Add(2)的需求吗？


## func (wg *WaitGroup) Done()

参数，返回值：

-	无

功能：

-	对WaitGroup的counter减1。

## func (wg *WaitGroup) Wait()

参数，返回值：

-	无

功能：

-	阻塞，直到WaitGroup中的所以过程完成。


代码示例：

	package main

	import (
    	"fmt"
    	"sync"
	)

	func wgProcess(wg *sync.WaitGroup,id int){
    	fmt.Printf("process:%d is going!\n",id)
    	wg.Done()
	}

	func main() {
    	wg := new(sync.WaitGroup)
    	for i := 0; i < 3; i++ {
        	wg.Add(1)
        	go wgProcess(wg, i)
    	}
    	wg.Wait()
	}


