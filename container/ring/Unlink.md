# func (r *Ring) Unlink(n int) *Ring

参数列表：

- `n`：要被移除的节点的数个

功能说明：

从节点`r`的下一个节点（包含该节点）开始移除`n % r.Len()`个节点。如果`n % r.Len() == 0`，则链表不会有改变。`r`不能为空。

代码实例：

```go

	package main

	import (
		"container/ring"
		"fmt"
	)

	func makeN(n int, begin int) *ring.Ring {
		r := ring.New(n)
		for i := begin; i < n+begin; i++ {
			r.Value = i
			r = r.Next()
		}
		return r
	}

	func main() {
		r := makeN(10, 0)
		r.Do(func(v interface{}) {
			fmt.Printf("%d ", v) // 输出：0 1 2 3 4 5 6 7 8 9 
		})
		fmt.Println()

		fmt.Printf("%d\n", r.Value) // 输出：0
		r.Unlink(5)
		r.Do(func(v interface{}){
			fmt.Printf("%d ", v)
		})
		// 输出：0 6 7 8 9 
	}

```