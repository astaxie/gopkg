## func copy(dst, src []Type) int

参数列表：

- dst []Type 与 src 同类型
- src []Type

返回值：

- int 返回被复制的元素数量

功能说明：

copy 内建函数将元素从来源切片复制到目标切片中。 （特殊情况是，它也能将字节从字符串复制到字节切片中）。来源和目标可以重叠。 copy 返回被复制的元素数量，它会是 len(src) 和 len(dst) 中较小的那个。

代码实例：

```go
package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 3)
	len := copy(dst, src[4:])
	fmt.Println(len, dst)
	len = copy(dst, src[0:])
	fmt.Println(len, dst)
	len = copy(dst, src)
	fmt.Println(len, dst)
}
```

输出：

~~~
1 [5 0 0]
3 [1 2 3]
3 [1 2 3]
~~~