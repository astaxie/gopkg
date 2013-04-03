## func close(c chan<- Type)

参数列表：

- c 信道

返回值：

- 无

功能说明：

close 内建函数关闭信道，该信道必须为双向的或只发送的。 它只能由发送者执行，不应由接收者执行，其效果是在最后发送的值被接收后停止该信道。 已关闭的信道 c 中最后一个值被接收后，任何从信道 c 的接收操作都会无阻塞成功， 关闭的信道会返回该信道元素类型的零值。对于已关闭的信道，形式

```go
x, ok := <-c
```

还会将 ok 置为 false。

代码实例：

```go
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)	// 因为使用了range调用,必须 close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {	// range 会不断从 channel 接收值，直到它被关闭。
		fmt.Println(i)
	}
}
```

输出:

~~~
0
1
1
2
3
5
8
13
21
34
~~~
