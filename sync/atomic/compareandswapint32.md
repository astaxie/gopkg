# func CompareAndSwapInt32(val *int32, old, new int32) (swapped bool)

参数：

-	val：需要进行修改的变量。
-	old：与*val进行比较的变量。
-	new：希望将val的值修改为new的值。

返回值：

-	swapped：如果*val=new成功则返回true，否则返回flase。

功能：

-	提供原子的compare-swap操作。过程是：
<pre><code>
	if *val == old {
		*val = new
		return true
	}
	return false
</code></pre>

兄弟函数：

-	fuc CompareAndSwapInt64
-	fuc CompareAndSwapPointer
-	fuc CompareAndSwapUint32
-	fuc CompareAndSwapUint64
-	fuc CompareAndSwapptr

在使用上只用参数类型的区别

代码示例：

	package main


	import (
    	"fmt"
    	"runtime"
    	"time"
    	"sync/atomic"
	)

	func c32To42(total *int32,id int) {
    	if atomic.CompareAndSwapInt32(total,32,42) {
        	fmt.Printf("32 to 42:id%d, work\n",id)
    	} else {
        	fmt.Printf("32 to 42:id%d, fail\n",id)
    	}
	}

	func c42To32(total *int32,id int) {
    	if atomic.CompareAndSwapInt32(total,42,32) {
        	fmt.Printf("42 to 32:id%d, work\n",id)
    	} else {
        	fmt.Printf("42 to 32:id%d, fail\n",id)
    	}
	}

	func main() {

    	runtime.GOMAXPROCS(2)		//使用多个处理器，不然都是顺序执行。

    	count := int32(32)

    	for i:=0; i<3; i++ {
       		go c32To42(&count,i)
       		go c42To32(&count,i)
    	}

    	time.Sleep(time.Second*1)
	}
