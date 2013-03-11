# heap包

## 概述:

任何实现了heap.Interface接口的对象都可以使用heap包提供的方法对堆进行操作(堆是一个完全二叉树)。通过对heap.Interface中的Less方法的不同实现，来实现最大堆和最小堆。通常堆的数据结构为一个一维数组。

## 函数列表:

- func Init(h Interface)
- func Pop(h Interface) interface{}
- func Push(h Interface, x interface{})
- func Remove(h Interface, i int) interface{}
- type Interface