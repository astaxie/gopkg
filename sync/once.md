#Once

Once结构如下：

	type Once struct {
		m		Mutex
		done	uint32
	}

##成员m

用于保证Once不会并发执行动作。

##成员done

用于保证Once只会执行一次。当done==1表明动作已经完成。

##func (o *Once) Do(f func())

参数：

-	func 函数类型

返回值：

-	无

功能：

-	执行且仅执行f()一次

代码示例：

	package main

	import (
		"fmt"
		"sync"
	)

	func main() {
		once := new(sync.Once)
		ch := make(chan int, 3)

		for i := 0; i < 3; i++ {
			go func(x int) {
				once.Do(func() {
					fmt.Printf("once %d\n", x)
				})
				fmt.Printf("%d\n", x)
				ch <- 1
			}(i)
		}

		for i := 0; i < 3; i++ {
			<-ch
		}
	}

输出：

once 0

0

1

2






