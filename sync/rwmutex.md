#RWMutex

RWMutex结构如下：

	type RWMutex struct {
		w			Mutex
		writerSem	uint32
		readerSem	uint32
		readerCount	int32
		readerWait	int32
	}

##成员

-	w			： 用于写锁，拒绝其他写操作。
-	writerSem	： 写锁信号量，用于阻塞或唤醒写锁争夺锁的行为。
-	readerSem	： 读锁信号量，写锁来时唤醒要进来的读锁。
-	readerCount	： 读锁的持有数。
-	readerWait	： 写锁需要等待的读操作数，最后一个读离开唤醒争用写锁的过程。

##func (rw *RWMutex) Rlock()

参数，返回值：

-	无

功能描述：

-	获取读锁，当之前以有一个写锁存在，阻塞。

##func (rw *RWMutex) RUlock()

参数，返回值：

-	无

功能描述：

-	放弃读锁，当存在有尝试进入的写锁且当前是最后一个读锁时，唤醒写锁争用过程。

##func (rw *RWMutex) Lock()

参数，返回值：

-	无

功能描述：

-	获取写锁，当之前存在读锁，阻塞。

##func (rw *RWMutex) Unlock()

参数，返回值：

-	无

功能描述：

-	放弃写锁，当存在有读锁争用过程被阻塞，唤醒所以读锁。

代码示例：


	package main


	import (
    	"fmt"
    	"runtime"
    	"sync"
	)

	func clickWithMutex(total *int,m *sync.RWMutex, ch chan int) {
    	for i := 0; i < 1000; i++ {
        	m.Lock()
        	*total += 1
        	m.Unlock()

        	if i==500 {
            	m.RLock()
            	fmt.Println(*total)
            	m.RUnlock()
        	}
    	}
    	ch <- 1
	}


	func main() {

    	runtime.GOMAXPROCS(2)		//使用多个处理器，不然都是顺序执行。

    	m := new(sync.RWMutex)
    	count := 0;

    	ch := make(chan int, 10)		//保证输出时count完了

    	for i := 0; i < 10; i++ {
        	go clickWithMutex(&count, m, ch)
    	}

    	for i := 0; i < 10; i++ {
        	<-ch
    	}

    	fmt.Printf("count:%d\n", count)
	}

>>>>>>> 4e6bb8a255a918bf287959e4f39c14c076f7cd1b








