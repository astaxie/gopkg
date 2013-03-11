# func Remove(h Interface, i int) interface{}

参数列表：

- h 实现了heap.Interface的堆对象
- i 将被移除的元素在堆中的索引号

返回值：

- interface{} 堆顶的元素

功能说明：

把索引号为i的元素从堆中移除。该方法的时间复杂度为O(log(n))，n为堆中元素的总和。

代码实例：
