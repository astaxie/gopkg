# func (l *List) PushFrontList(ol *List)

参数列表：

- `ol`：将被插入到链表`l`开头的链表

功能说明：

把一个链表存到链表开头

代码实例：

```go

	package main

	import (
		"fmt"
		"container/list"
	)

	func main(){
		l1 := list.New()
		l1.PushFront("a")
		l1.PushFront("b")
		fmt.Println(l1.Len()) // 输出：2
		
		l2 := list.New()
		l2.PushFront("c")
		l2.PushFront("d")
		
		l1.PushFrontList(l2) // l2中所有节点的list字段都是l2，在l2的所有节点都加到l1的开头后，list字段编程了l1
		fmt.Println(l1.Len()) // 输出：4
		elementFromL1 := l1.Front();
		fmt.Println(elementFromL1.Value) // 输出：d
		l1.Remove(elementFromL1) // 删除开头的元素
		fmt.Println(l1.Len()) // 输出：3
		fmt.Println(l1.Front().Value) // 输出：c
	}

```