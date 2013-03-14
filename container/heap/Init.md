# func Init(h Interface)

参数列表：

- h 实现了heap.Interface的堆对象

功能说明：

在对堆h进行操作前必须保证堆已经初始化（即符合堆结构），该方法可以在堆中元素的顺序不符合堆的要求时调用，调用后堆会调整为标准的堆结构，该方法的时间复杂度为：O(n)，n为堆h中元素的总个数。

代码实例：

	package main
	
	import (
		"container/heap"
		"fmt"
	)
	
	type myHeap []int  // 定义一个堆，存储结构为数组
	
	// 实现了heap.Interface中组合的sort.Interface接口的Less方法，
	// 此处的实现是元素i小于元素j，所以此堆是最小堆（任意子树的根元素是子树中最小的）。
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
	
	// 按层来遍历和打印堆数据
	func (h myHeap) printHeap() {
		n := 1
		levelCount := 1
		for n <= h.Len() {
			fmt.Println(h[n-1 : n-1+levelCount])
			n += levelCount
			levelCount *= 2
		}
	}
	
	func main() {
		data := [7]int{13, 25, 1, 9, 5, 12, 11}
		aheap := new(myHeap)
		// 将数组中的元素依次存入堆中
		for _, value := range data {
			aheap.Push(value)
		}
	
		// 此时堆数组内容为：13, 25, 1, 9, 5, 12, 11
		aheap.printHeap()
		// Output：
		//  [13]
		//  [25 1]
		//  [9 5 12 11]
	
		heap.Init(aheap) // 对堆进行初始化
	
		// 此时堆数组内容为：1,5,9,11,12,13,25	
		aheap.printHeap()
		// Output：
		//	[1]
		//	[5 11]
		//	[9 25 12 13]
	}
	