# func (r *Ring) Prev() *Ring

返回值：

- `*Ring`：指向上一个节点的指针

功能说明：

获得指向上一个节点的指针

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
		r := makeN(5, 0)
		r.Do(func(v interface{}) {
			fmt.Printf("%d ", v) // 输出：0 1 2 3 4 
		})
		fmt.Println()

		fmt.Printf("%d ", r.Value)
		rr := r.Prev()
		for ; rr != r; rr = rr.Prev() {
			fmt.Printf("%d ", rr.Value)
		}
		// 输出：0 1 2 3 4 
	}

```
