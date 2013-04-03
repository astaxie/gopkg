# func (r *Ring) Move(n int) *Ring

参数列表：

- `n`：指针`r`在双向链表上移动的位置的个数。n>0时，为正向移动；反之为反向移动。

返回值：

- *Ring：移动结束后，指针`r`指向的节点

功能说明：

指向节点`r`的指针，正向或者逆向移动`n % r.Len()`个节点，并返回这个指针移动后指向的节点。但是`r.Move(n)`不对改变r的值，`r`不能为空

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

		fmt.Println(r.Value) // 输出：0
		rr := r.Move(3)
		fmt.Println(r.Value) // 输出：0
		fmt.Println(rr.Value) // 输出：3
	}

```