#Cond

Cond结构如下：

	type Cond struct {
		L			Locker
		m			Mutex
		oldWaiters	int
		oldSema		*uint32
		newWaiters	int
		newSema		*uint32
	}
	
##成员

-	L：外部传入的Locker实现，用于保护Cond的操作。
-	m：用于内部结构的保护。
-	oldWaiters：上一代的waiter计数。
-	oldSema：上一代的信号量。
-	newWaiters：新一代的waiter计数。
-	newSema：新一代的信号量。

ps：分代管理是为了在使用`Signal()`唤醒一个waiter时，优先唤醒oldWaiter。

## func NewCond(l Locker) *Cond

参数：

-	Locker接口，可以自己实现。

返回值：

-	Cond指针。

功能：

-	创建Cond。

## func (c *Cond) Wait()

参数，返回值：

-	无

功能：

-	添加waiter，使用时要注意先调用c.L.Lock()。详见代码示例。

## func (c *Cond) Signal()

参数，返回值：

-	无

功能：

-	唤醒一个waiter。优先唤醒上一代的waiter，如果oldWaiters==0，新一代变成上一代。详见源码cond.go

## func (c *Cond) Broadcast()

参数，返回值：

-	无

功能：

-	唤醒所有waiter，包括上一代和新一代。


代码示例：


	package main

	import (
    	"fmt"
    	"time"
    	"sync"
	)

	func waiter(cond *sync.Cond,id int) {
    	cond.L.Lock()
    	cond.Wait()
    	cond.L.Unlock()
    	fmt.Printf("Waiter:%d wake up!\n",id)
	}

	func main() {
    	locker := new(sync.Mutex)
    	cond := sync.NewCond(locker)  //使用Mutex作为Locker
    
    	for i := 0; i < 3; i++ {		//生成waiter
        	go waiter(cond,i)
    	}
    	time.Sleep(time.Second * 1)		//等待waiter到位

    	cond.L.Lock()
    	cond.Signal()					//唤醒一个waiter
    	cond.L.Unlock()

    	for i := 3; i < 5; i++ {		//生成新一代waiter
        	go waiter(cond,i)
    	}
    	time.Sleep(time.Second * 1)

    	cond.L.Lock()
    	cond.Signal()					//唤醒的将是上一代（id<3）的waiter之一
    	cond.L.Unlock()

    	cond.L.Lock()
    	cond.Broadcast()				//唤醒所以waiter
    	cond.L.Unlock()
    	time.Sleep(time.Second * 1)
	}
>>>>>>> 4e6bb8a255a918bf287959e4f39c14c076f7cd1b
