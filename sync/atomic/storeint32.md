#func StoreInt32(addr *int32, val int32)

参数：

-	addr：需要操作的变量地址。
-	val： 要存进addr里的值。

返回值：

-	无

兄弟函数：

-	func StoreInt64
-	func StorePointer
-	func StoreUint32
-	func StoreUint64
-	func StoreUintptr

在使用上的区别只要操作类型不一样。

功能说明：

-	提供互斥写操作。

代码示例：

	package main


	import "fmt"
	import "time"
	import "sync/atomic"
	import "runtime"

	func read(num *int32) {
    	for {
        	if *num==32 {
            	fmt.Println("ops:", 32)
        	}
        	if *num==42 {
            	fmt.Println("ops:", 42)
        	}
        	time.Sleep(time.Nanosecond)
    	}
	}

	func main() {

    	runtime.GOMAXPROCS(4)

    	var ops int32 = 0

    	for i := 0; i < 4; i++ {
        	go read(&ops)
    	}

    	atomic.StoreInt32(&ops, 32)

    	time.Sleep(time.Nanosecond*5)

    	ops = 42					//42的输出会不稳定。
    	fmt.Println("changed")

    	time.Sleep(time.Nanosecond*20)
	}

>>>>>>> 4e6bb8a255a918bf287959e4f39c14c076f7cd1b
