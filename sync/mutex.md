#Mutex

Mutex是对Locker的一个实现。结构如下：

	type Mutex struct {
		state int32
		sema  uint32
	}
	
##state
成员state用表明当前锁是处于被占用（state==1）还是空闲（state==0）。使用`atomic.CompareAndSwapInt32()`进行修改。



##sema
当前锁占用失败，系统监听成员sema。当sema==1，表明锁被释放。当前争用的的过程将被系统唤醒尝试去获取锁。当sema==0，当前争用过程会被系统暂停。



##func Lock() && func Unlock()

参数，返回值：

- 无
	
功能说明：

- 获取和释放当前锁。

代码示例：


	package main


	import (
    	"fmt"
    	"runtime"
    	"sync"
	)

	func click(total *int,ch chan int) {
    	for i := 0; i < 1000; i++ {
        	*total += 1
    	}
    	ch <- 1
	}

	func clickWithMutex(total *int,m *sync.Mutex, ch chan int) {
    	for i := 0; i < 1000; i++ {
        	m.Lock()
        	*total += 1
        	m.Unlock()
    	}
    	ch <- 1
	}


	func main() {

    	runtime.GOMAXPROCS(2)		//使用多个处理器，不然都是顺序执行。

    	m := new(sync.Mutex)
    	count1 := 0;
    	count2 := 0;

    	ch := make(chan int, 200)		//保证输出时count完了

    	for i := 0; i < 100; i++ {
        	go click(&count1, ch)
    	}
    	for i := 0; i < 100; i++ {
        	go clickWithMutex(&count2, m, ch)
    	}

    	for i := 0; i < 200; i++ {
        	<-ch
    	}

    	fmt.Printf("count1:%d\ncount2:%d\n", count1,count2)
	}
>>>>>>> 4e6bb8a255a918bf287959e4f39c14c076f7cd1b

</code></pre>

程序输出：

- count1：55523		//不定
- coutn2：100000