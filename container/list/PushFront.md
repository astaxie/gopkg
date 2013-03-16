# func (l *List) PushFront(value interface{}) *Element

参数列表：

- `value`：将被存到链表开头的任意对象

返回值：

- `*Element`：被存到开头的节点的指针

功能说明：

把一个对象存到链表开头，并返回这个节点

代码实例：

```go

	package main

	import (
		"fmt"
		"container/list"
	)

	func main(){
		l := list.New()
		l.PushFront("a")
		fmt.Println(l.Front().Value) // 输出：a
		l.PushFront("b")
		fmt.Println(l.Front().Value) // 输出：b
		l.PushFront("c")
		fmt.Println(l.Front().Value) // 输出：c
	}

```