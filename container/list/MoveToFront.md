# func (l *List) MoveToFront(e *Element)

参数列表：

- `e`：链表中的节点

功能说明：

把节点`e`移到链表的开头

代码实例：

```go

	package main

	import (
		"container/list"
		"fmt"
	)

	func main() {
		l := list.New()
		l.PushBack("a")
		l.PushBack("b")
		e := l.PushBack("c")
		
		fmt.Println(l.Front().Value) // 输出：a
		fmt.Println(l.Back().Value)  // 输出：c
		
		l.MoveToFront(e)
		
		fmt.Println(l.Front().Value) // 输出：c
		fmt.Println(l.Back().Value)  // 输出：b
	}

```