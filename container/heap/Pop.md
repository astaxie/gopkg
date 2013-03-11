# func Pop(h Interface) interface{}

参数列表：

- h 实现了heap.Interface的堆对象

返回值：

- interface{} 堆顶的元素

功能说明：

从堆h中取出堆顶的元素。根据h的Less方法实现的不同，堆顶元素可以是最大的元素或者是最小的元素。该方法的时间复杂度为O(log(n))，n为堆中元素的总和。

代码实例：
