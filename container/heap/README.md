# heap包

## 概述:

任何实现了`heap.Interface`接口的对象都可以使用heap包提供的方法对堆进行操作(堆是一个完全二叉树)。通过对`heap.Interface`中的`Less`方法的不同实现，来实现最大堆和最小堆。通常堆的数据结构为一个一维数组。

维基百科：[堆（数据结构）](http://zh.wikipedia.org/wiki/%E5%A0%86_(%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84)) "堆（数据结构）"

## 函数列表:

- [func Init(h Interface)](Init.md)
- [func Pop(h Interface) interface{}](Pop.md)
- [func Push(h Interface, x interface{})](Push.md)
- [func Remove(h Interface, i int) interface{}](Remove.md)
- [type Interface](Interface.md)