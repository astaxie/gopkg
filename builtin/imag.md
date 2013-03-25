## func imag(c ComplexType) FloatType

参数列表：

- c ComplexType 复数

返回值：

- FloatType 复数 c 的虚部

功能说明：

imag 内建函数返回复数 c 的虚部。 返回值类型与 c 的浮点类型对应。

代码实例：

```go
package main

import "fmt"

func main() {
	fmt.Println(imag(10 + 2i))
}
```

输出：

~~~
2
~~~