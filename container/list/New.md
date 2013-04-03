# func New() *List

返回值：

- `*List`：空链表的指针

功能说明:

创建一个空链表，链表的长度为0，开头和末尾节点都是`nil`

代码实例：

```go

	package main

	import (
		"container/list"
		"fmt"
	)

	func main() {
		l := list.New()
		e := l.Front()        // 取出链表开头的节点，即节点a
		fmt.Println(e == nil) // 输出：true
	}

```