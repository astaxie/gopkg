# func (l *List) Front() *Element

返回值：

- `*Element`：链表中第一个节点的指针，如果链表长度为0，则为`nil`

功能说明:

获得第一个节点的指针，如果链表长度为0，则为`nil`

代码实例：

```go

	package main

	import (
		"container/list"
		"fmt"
	)

	func main() {
		l := list.New()
		e := l.Front()
		fmt.Println(e == nil)
		l.PushFront("b")
		e = l.Front()
		fmt.Println(e.Value) // 输出：b
	}

```