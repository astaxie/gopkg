# func (r *Ring) Do(f func(interface{}))

参数列表：

- `f`：一个回调函数，该函数的参数为环形双向链表中的节点的`Value`字段值

功能说明：

正向遍历环形双向链表，并对每个链表中的元素执行回调函数`f`，如果这个回调函数`f`会修改链表`r`，那这个回调函数的行为将不可确定

代码实例：

```go

	package main

	import (
		"container/ring"
		"fmt"
	)

	func main() {
		r := ring.New(10)
		n := 0
		r.Do(func(v interface{}){
			n++
			v = n
			fmt.Printf("%d ", v) // 输出：1 2 3 4 5 6 7 8 9 10 
		})
	}

```