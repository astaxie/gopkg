# func (r *Ring) Len() int

返回值：

- `int`：环形链表中元素的个数

功能说明:

遍历环形双向链表来统计链表中的元素的个数

代码实例：

```go

	package main

	import (
		"container/ring"
		"fmt"
	)

	func main() {
		r := ring.New(10)
		fmt.Println(r.Len() == 10) // 输出：true
	}

```