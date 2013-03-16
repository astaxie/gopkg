# type Interface

接口代码：

```go

	type Interface interface {
	    sort.Interface
	    Push(x interface{})
	    Pop() interface{}
	}
	
```

功能说明：

这是堆的接口，`heap`包里面的方法只是提供堆的一些堆算法操作，要想使用这些算法操作，就必须实现这些接口，每个接口方法都有具体的含义，堆本身的数据结构由这个接口的具体实现决定，可以是数组、列表。

接口方法：

- sort.Interface 

	参见[sort.Interface](../../sort/README.md)接口。
	
	其中的Less(i, j int)方法的实现决定了堆是最大堆还是最小堆，如果该方法的实现是指元素i是否小于元素j，则是最小堆，反之为最大堆，代码实例：

```go
		type myHeap []int  // 定义一个堆，存储结构为数组
		
		// 最小堆的Less方法实现
		func (h *myHeap) Less(i, j int) bool {
			return (*h)[i] < (*h)[j]
		}
		
		// 最大堆的Less方法实现
		func (h *myHeap) Less(i, j int) bool {
			return (*h)[i] < (*h)[j]
		}
```

- Push(x interface{}) 

	参数列表：

	x 将存到堆中的元素

	功能说明：

	把元素x存放到索引号为Len()的位置上，比如，一个列表中元素有7个，索引号从0开始，那么x将被存放到索引号为7的位置上，即最末尾。
	
	代码实例：

```go
		type myHeap []int  // 定义一个堆，存储结构为数组
		
		func (h *myHeap) Push(v interface{}) {
			*h = append(*h, v.(int))
		}
```

- Pop() interface{}
	
	返回值：

	被移除的索引号为Len()-1的元素

	功能说明：
	
	把索引号为Len()-1的元素移除并返回这个被移除的元素。比如，如果这个堆是一个数组，那么才操作就是把数组的最后一个元素移除并返回。

	代码实例：

```go
		type myHeap []int  // 定义一个堆，存储结构为数组
		
		func (h *myHeap) Pop() (v interface{}) {
			*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
			return
		}
```
