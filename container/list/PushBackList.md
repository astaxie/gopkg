# func (l *List) PushBackList(ol *List)

参数列表：

- `ol`：将被插入到链表`l`末尾的链表

功能说明：

把一个链表存到链表末尾

代码实例：

```go

	package main

	import (
		"fmt"
		"container/list"
	)

	func main(){
		l1 := list.New()
		l1.PushBack("a")
		l1.PushBack("b")
		fmt.Println(l1.Len()) // 输出：2
		
		l2 := list.New()
		l2.PushBack("c")
		l2.PushBack("d")
		
		l1.PushBackList(l2) // l2中所有节点的list字段都是l2，在l2的所有节点都加到l1的末尾后，list字段编程了l1
		fmt.Println(l1.Len()) // 输出：4
		elementFromL1 := l1.Back();
		fmt.Println(elementFromL1.Value) // 输出：d
		l1.Remove(elementFromL1) // 删除末尾的元素
		fmt.Println(l1.Len()) // 输出：3
		fmt.Println(l1.Back().Value) // 输出：c
	}

```