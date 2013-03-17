# func (l *List) Remove(e *Element) interface{}

参数列表：

- `e`：将被删除的节点，该节点必须是属于链表`l`的

返回值：

- `interface{}`：被删除的节点的内容

功能说明：

删除指定的节点，并返回这个节点的内容

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
		elementFromL2 := l2.PushBack("d")
		
		l1.PushBackList(l2) // l2中所有节点的list字段都是l2，在l2的所有节点都加到l1的末尾后，list字段编程了l1
		fmt.Println(l1.Len()) // 输出：4
		
		elementFromL1 := l1.Back();
		fmt.Println(elementFromL1.Value) // 输出：d
		l1.Remove(elementFromL2) // elementFromL2是属于l2的，不属于l1，所以l1不会有节点被删除
		fmt.Println(l1.Len()) // 输出：4
		fmt.Println(l1.Back().Value) // 输出：d
		l1.Remove(elementFromL1) // 成功的删除
		fmt.Println(l1.Len()) // 输出：3
		fmt.Println(l1.Back().Value) // 输出：c
	}

```