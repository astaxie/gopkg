# func (l *List) InsertBefore(value interface{}, mark *Element) *Element

参数列表：

- `value`：要插入的数据的内容
- `mark`：链表中的一个节点指针

返回值：

- `*Element`：被插入的节点指针，该节点的`Value`为数据内容

功能说明：

把数据`value`插入到`mark`节点的前面，并返回这个被插入的节点。

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
		e := l.PushBack("c")
		fmt.Println(e.Prev().Value) // 输出：a
		l.InsertBefore("b", e) // 在c前面插入b
		fmt.Println(e.Prev().Value) // 输出：b
	}

```