# func Push(h Interface, x interface{})

参数列表：

- h 实现了heap.Interface的堆对象
- x 将被存到堆中的元素对象

功能说明：

把元素x存到堆中。该方法的时间复杂度为O(log(n))，n为堆中元素的总和。

代码实例：
