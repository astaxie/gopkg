#func LoadInt32(addr *int32) (val int32)

参数：

-	addr：读取的地址。

返回值

-	val：读取的值。

功能：

-	提供互斥读，也就是说读期间无法进行其它读写操作。

兄弟函数：

-	func LoadInt64
-	func LoadPointer
-	func LoadUint32
-	func LoadUint64
-	func LoadUintptr

在使用上只要类型的区别。

代码示例:


	package main


	import "fmt"
	import "time"
	import "sync/atomic"
	import "runtime"

	func main() {

    	runtime.GOMAXPROCS(2)

    	var ops int32 = 0

    	for i := 0; i < 2; i++ {
        	go func() {
            	for i:=0; i<100; i++{
                	time.Sleep(time.Nanosecond)
                	atomic.AddInt32(&ops, 1)
                	if i==50 {
                   		fmt.Println("ops:", atomic.LoadInt32(&ops),ops)
                	}
            	}
        	}()
    	}

    	time.Sleep(time.Second)
	}
