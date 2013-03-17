# func (r *Ring) Link(s *Ring) *Ring

参数列表：

- `s`：环形双向链表

返回值：

- `*Ring`：`s`和`r`相连前`r.Next()`的值，也是相连后`s.Next()`的值

功能说明:

把一个环形双向链表`s`与环形双向链表`r`相链接，链接后`r.Next()`为s，并返回相连前时`r.Next()`的值。`r`不能为空。

- 如果`s`和`r`不是同一个环形链表，则相连后，只产生一个环形链表，并返回相连前时`r.Next()`的值，也是相连后`s.Next()`的值。

- 如果`s`和`r`是同一个环形链表，但`s != r`时，相连后，产生两个环形链表，其中一个是由`r`和`s`之间的节点构成（不包括`r`和`s`），返回值为相连前时`r.Next()`的值，即`r`和`s`之间的节点（不包括`r`和`s`）构成的环形链表的表头节点。

- 如果`s`和`r`是同一个环形链表，且`s == r`时，相连后，产生两个环形链表，其中一个是由`r`指向的节点构成的长度为1的环形链表，其他节点构成另一个环形链表，返回值为相连前时`r.Next()`的值，即其他节点构成的环形链表的表头节点。

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

	// Link两个不同的环形链表
	func linkDiffRing() {
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

	// 如果两个链表是同一个链表，但是两个指针不是指这个链表中的节点时，即s != r时
	func linkSameRing() {
		r1 := makeN(10, 0)
		r1.Do(func(v interface{}) {
			fmt.Printf("%d ", v) // 输出：0 1 2 3 4 5 6 7 8 9 
		})
		fmt.Println()

		r2 := r1.Move(5)
		fmt.Println(r1.Value, r2.Value) // 输出：0 5
		
		r3 := r1.Link(r2)
		
		r1.Do(func(v interface{}) {
			fmt.Printf("%d ", v.(int)) // 输出：0 5 6 7 8 9 
		})
		fmt.Println()
		r3.Do(func(v interface{}) {
			fmt.Printf("%d ", v.(int)) // 输出：1 2 3 4 
		})
	}

	// 如果两个链表是同一个链表，且两个指针指向这个链表中的同一个节点时，即s == r时
	func linkSameElement() {
		r1 := makeN(10, 0)
		r1.Do(func(v interface{}) {
			fmt.Printf("%d ", v) // 输出：0 1 2 3 4 5 6 7 8 9 
		})
		fmt.Println()
		fmt.Println(r1.Value) // 输出：0
		
		r2 := r1.Link(r1)
		
		r1.Do(func(v interface{}) {
			fmt.Printf("%d ", v.(int)) // 输出：0 
		})
		fmt.Println()
		r2.Do(func(v interface{}) {
			fmt.Printf("%d ", v.(int)) // 输出：1 2 3 4 5 6 7 8 9 
		})
	}

	func main() {
		linkDiffRing()
		fmt.Println()
		linkSameRing()
		fmt.Println()
		linkSameElement()
	}

```
