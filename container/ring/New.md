# func New(n int) *Ring

参数列表：

- `n`：环形双向链表的节点的个数

返回值：

- `*Ring`：空链表的指针

功能说明:

创建一个有n个节点的环形双向链表

代码实例：

```go

	package main

	import (
		"container/ring"
		"fmt"
	)

	func main() {
		r := ring.New(10)
		fmt.Println(r.Len() == 10) // 输出：true
	}

```