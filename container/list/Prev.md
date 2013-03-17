# func (e *Element) Prev() *Element

返回值：

- `*Element`：链表中该节点的上一个节点元素的指针，如果该节点是第一个节点，则返回`nil`

功能说明：

获得该节点在链表中的上一个节点元素的指针，如果该节点是第一个节点，则返回`nil`

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
		e := l.Back()   // 取出链表开头的节点，即节点b
		if e != nil {
			fmt.Println(e.Prev().Value) // 输出：a
		}
	}

```