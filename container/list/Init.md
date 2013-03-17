# func (l *List) Init() *List

返回值：

- `*List`：初始化或者清空后的链表

功能说明:

初始化或者清空链表，该方法调用后，链表的长度为0

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
		l.PushBack("c")
		fmt.Println(l.Len()) // 输出：3
		
		l.Init()
		fmt.Println(l.Len()) // 输出：0
	}

```