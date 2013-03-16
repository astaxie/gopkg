# func (l *List) PushBack(value interface{}) *Element

参数列表：

- `value`：将被存到链表末尾的任意对象

返回值：

- `*Element`：被存到末尾的节点的指针

功能说明：

把一个对象存到链表末尾，并返回这个节点

代码实例：

```go

	package main

	import (
		"fmt"
		"container/list"
	)

	func main(){
		l := list.New()
		l.PushBack("a")
		fmt.Println(l.Back().Value) // 输出：a
		l.PushBack("b")
		fmt.Println(l.Back().Value) // 输出：b
		l.PushBack("c")
		fmt.Println(l.Back().Value) // 输出：c
	}

```