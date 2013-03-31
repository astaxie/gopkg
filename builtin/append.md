## func append(slice []Type, elems ...Type) []Type

参数列表：

- slice Type类型切片
- elems Type类型元素
- ... 更多的Type类型元素

返回值：

- Type类型切片

功能说明：

append 内建函数将元素追加到切片的末尾。 若它有足够的容量，其目标就会重新切片以容纳新的元素。否则，就会分配一个新的基本数组。 append 返回更新后的切片。因此必须存储追加后的结果，通常为包含该切片自身的变量：

```go
slice = append(slice, elem1, elem2)
slice = append(slice, anotherSlice...)
```

代码实例：

```go
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	s := a[1:3]
	c := append(s, 6)
	fmt.Println("a ==", a[:cap(a)])
	fmt.Println("s ==", s[:cap(s)])
	fmt.Println("c ==", c[:cap(c)])

	d := append(a, 6)
	fmt.Println("a ==", a[:cap(a)])
	fmt.Println("d ==", d[:cap(d)])
}
```

输出:
~~~
a == [1 2 3 6 5]
s == [2 3 6 5]
c == [2 3 6 5]
a == [1 2 3 6 5]
d == [1 2 3 6 5 6 0 0 0 0]
~~~