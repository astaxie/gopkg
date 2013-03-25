## func complex(r, i FloatType) ComplexType

参数列表：

- r FloatType 实部
- i FloatType 虚部

返回值：

- ComplexType

功能说明：

complex 内建函数将两个浮点数值构造成一个复数值。 其实部和虚部的大小必须相同，即 float32 或 float64（或可赋予它们的），其返回值 即为对应的复数类型（complex64 对应 float32，complex128 对应 float64）。

代码实例：

```go
package main

import "fmt"

func main() {
	fmt.Println(complex(10, 2))
}
```

输出：

~~~
(10+2i)
~~~