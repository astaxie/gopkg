# func (e *Element) Next() *Element

返回值：

- `*Element`：链表中该节点的下一个节点元素的指针，如果该节点是最后一个节点，则返回`nil`

功能说明：

获得该节点在链表中的下一个节点元素的指针，如果该节点是最后一个节点，则返回`nil`

代码实例：

```go

	package main

	import (
		"container/list"
		"fmt"
	)

	func main() {
		l := list.New()
		l.PushBack("a") // 把a加入链表的末尾
		l.PushBack("b") // 把b加入链表的末尾
		e := l.Front()   // 取出链表开头的节点，即节点a
		if e != nil {
			fmt.Println(e.Next().Value) // 输出：b
		}
	}

```