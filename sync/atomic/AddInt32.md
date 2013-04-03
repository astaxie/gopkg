# fuc AddInt32(val *int32, delta int32) (new int32)


参数：

-	val：即将要操作的变量。
-	delta：增量。

返回值：

-	new：返回*val+delta。


功能：

-	在对val有竞争的时候，确保`*val=*val+delat`是一个原子操作。

兄弟函数：

-	func AddInt64
-	func AddUint32
-	func AddUint64
-	func AddUintptr

在使用上只要操作类型的区别。



代码示例：

	package main


	import (
    	"fmt"
    	"runtime"
    	"sync/atomic"
	)

	func click(total *int,ch chan int) {
    	for i := 0; i < 1000; i++ {
        	*total += 1
    	}
    	ch <- 1
	}

	func clickat(total *int32, ch chan int) {
    	for i := 0; i < 1000; i++ {
        	atomic.AddInt32(total, 1)
    	}
    	ch <- 1
	}


	func main() {

    	runtime.GOMAXPROCS(2)		//使用多个处理器，不然都是顺序执行。

    	count1 := 0;
    	count2 := int32(0)

    	ch := make(chan int, 200)		//保证输出时count完了

    	for i := 0; i < 100; i++ {
        	go click(&count1, ch)
    	}
    	for i := 0; i < 100; i++ {
        	go clickat(&count2, ch)
    	}

    	for i := 0; i < 200; i++ {
        	<-ch
    	}

    	fmt.Printf("count1:%d\ncount2:%d\n", count1,count2)
	}

