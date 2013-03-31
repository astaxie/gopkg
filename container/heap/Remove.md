# func Remove(h Interface, i int) interface{}

参数列表：

- `h`：实现了`heap.Interface`的堆对象
- `i`：将被移除的元素在堆中的索引号

返回值：

- interface{}：堆顶的元素

功能说明：

把索引号为i的元素从堆中移除。该方法的时间复杂度为O(log(n))，n为堆中元素的总和。

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
		i := 0          // 当前元素的索引号
		levelCount := 1 // 当前层级的元素个数
		for i+1 <= h.Len() {
			fmt.Println(h[i : i+levelCount])
			i += levelCount
			if (i + levelCount*2) <= h.Len() {
				levelCount *= 2
			} else {
				levelCount = h.Len() - i
			}
		}
		fmt.Println()
	}

	func main() {
		data := [7]int{13, 25, 1, 9, 5, 12, 11}
		aheap := new(myHeap) // 创建空堆

		// 用heap包中的Push方法将数组中的元素依次存入堆中，
		// 每次Push都会保证堆是规范的堆结构
		for _, value := range data {
			heap.Push(aheap, value)
		}
		fmt.Println(*aheap) // 输出：[1 5 11 25 9 13 12]
		aheap.printHeap()
		// 输出：
		//  [1]
		//  [5 11]
		//  [25 9 13 12]

		value := heap.Remove(aheap, 2) // 删除索引号为3的元素（即数组中的第4个元素）
		fmt.Println(value) // 输出：25
		aheap.printHeap()
		// 输出：
		//	[1]
		//	[5 12]
		//	[25 9 13]
	}
```
