# func (l *List) Back() *Element

返回值：

- `*Element`：链表中最后一个节点的指针，如果链表长度为0，则为`nil`

功能说明:

获得最后一个节点的指针，如果链表长度为0，则为`nil`

代码实例：

```go

	package main

	import (
		"container/list"
		"fmt"
	)

	func main() {
		l := list.New()
		e := l.Back()
		fmt.Println(e == nil)
		l.PushBack("a")
		e = l.Back()
		fmt.Println(e.Value) // 输出：a
	}

```