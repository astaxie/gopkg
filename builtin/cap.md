## func cap(v Type) int

参数列表：

- v Type

返回值：

- int

功能说明：

The cap built-in function returns the capacity of v, according to its type:
cap 内建函数返回 v 的容量，返回值取决 v 的具体类型：


Channel: the channel buffer capacity, in units of elements;
if v is nil, cap(v) is zero.

数组：v 中元素的数量（与 len(v) 相同）。
数组指针：*v 中元素的数量（与 len(v) 相同）。
切片：在重新切片时，切片能够达到的最大缓存长度；若 v 为 nil，len(v) 即为零。
信道：以信道元素数量为单位返回信道缓存容量；若 v 为 nil，len(v) 即为零。