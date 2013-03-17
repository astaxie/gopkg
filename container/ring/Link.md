# func (r *Ring) Link(s *Ring) *Ring

参数列表：

- `s`：环形双向链表

返回值：

- `*Ring`：`s`和`r`相连前`r.Next()`的值

功能说明:

把一个环形双向链表`s`与环形双向链表`r`相链接，链接后`r.Next()`为s，并返回链接前时`r.Next()`的值

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
		r1 := makeN(5, 0)
		r1.Do(func(v interface{}) {
			fmt.Printf("%d ", v) // 输出：0 1 2 3 4 
		})
		fmt.Println()

		r2 := makeN(5, 10)
		r2.Do(func(v interface{}) {
			fmt.Printf("%d ", v) // 输出：10 11 12 13 14 
		})
		fmt.Println()

		r1.Link(r2)
		fmt.Println(r1.Value) // 输出：0

		r1.Do(func(v interface{}) {
			fmt.Printf("%d ", v.(int)) // 输出：0 10 11 12 13 14 1 2 3 4 
		})
	}

```