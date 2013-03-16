# func (l *List) Len() int

返回值：

- `int`：链接中节点的个数

功能说明：

获得链接中节点的个数

代码实例：

```go

	package main

	import (
		"container/list"
		"fmt"
	)

	func main() {
		l := list.New()
		fmt.Println(l.Len()) // 输出：0
		l.PushBack("a")
		fmt.Println(l.Len()) // 输出：1
		e := l.PushBack("b")
		fmt.Println(l.Len()) // 输出：2
		l.PushBack("c")
		fmt.Println(l.Len()) // 输出：3
		l.Remove(e)
		fmt.Println(l.Len()) // 输出：2
	}

```