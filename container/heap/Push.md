# func Push(h Interface, x interface{})

参数列表：

- `h`：实现了`heap.Interface`的堆对象
- `x`：将被存到堆中的元素对象

功能说明：

把元素`x`存到堆中。该方法的时间复杂度为O(log(n))，n为堆中元素的总和。

代码实例：

```go
	package main

	import (
		"container/heap"
		"fmt"
	)

	type myHeap []int // 定义一个堆，存储结构为数组

	// 实现了heap.Interface中组合的sort.Interface接口的Less方法
	func (h *myHeap) Less(i, j int) bool {
		return (*h)[i] < (*h)[j]
	}

	// 实现了heap.Interface中组合的sort.Interface接口的Swap方法
	func (h *myHeap) Swap(i, j int) {
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
	}

	// 实现了heap.Interface中组合的sort.Interface接口的Push方法
	func (h *myHeap) Len() int {
		return len(*h)
	}

	// 实现了heap.Interface的Pop方法
	func (h *myHeap) Pop() (v interface{}) {
		*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
		return
	}

	// 实现了heap.Interface的Push方法
	func (h *myHeap) Push(v interface{}) {
		*h = append(*h, v.(int))
	}

	// 按层来遍历和打印堆数据，从堆顶开始
	func (h myHeap) printHeap() {
		i := 0 // 当前元素的索引号
		levelCount := 1 // 当前层级的元素个数
		for i+1 <= h.Len() {
			fmt.Println(h[i : i+levelCount])
			i += levelCount
			if (i + levelCount * 2) <= h.Len() {
				levelCount *= 2
			} else {
				levelCount = h.Len() - i
			}
		}
	}

	func main() {
		data := [7]int{13, 25, 1, 9, 5, 12, 11}
		aheap := new(myHeap) // 此时是空堆
		
		// 用heap包中的Push方法将数组中的元素依次存入堆中，
		// 每次Push都会保证堆是规范的堆结构
		for _, value := range data {
			heap.Push(aheap, value)
			fmt.Printf("插入%d\n", value)
			aheap.printHeap()
			fmt.Println()
		}
		fmt.Println(*aheap)
		// 输出：
		//	插入13
		//	[13]
		//	
		//	插入25
		//	[13]
		//	[25]
		//	
		//	插入1
		//	[1]
		//	[25 13]
		//	
		//	插入9
		//	[1]
		//	[9 13]
		//	[25]
		//	
		//	插入5
		//	[1]
		//	[5 13]
		//	[25 9]
		//	
		//	插入12
		//	[1]
		//	[5 12]
		//	[25 9 13]
		//	
		//	插入11
		//	[1]
		//	[5 11]
		//	[25 9 13 12]
		//	
		//	[1 5 11 25 9 13 12]

	}
```
