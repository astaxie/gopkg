## func panic(v interface{})

参数列表：

- v 接口

返回值：

- 无

功能说明：

panic 内建函数停止当前 goroutine 的正常执行。 当函数 F 调用 panic，函数 F 的执行被终止，并且 F 中的延迟函数会正常执行，然后 F 返回到调用它的地方。对于调用者 G，F 的行为就像调用了 panic，即终止 G 的执行并运行所有被推迟的函数。这一过程会到该程序中所有函数都按相反的顺序停止执行之后。此时，该程序会被终止，并产生错误报告， 包括引发该 panic 的实参值。此终止序列称为恐慌过程，可以通过内建函数 recover 控制。

代码实例：

```go
package main

import (
	"fmt"
)

func f() {
	defer func() {
		fmt.Println("f() defer")
	}()
	fmt.Println("f() before panic")
	panic(0)
	// 此后的代码不会被执行
	fmt.Println("after panic")
	defer func() {
		fmt.Println("f() defer agin")
	}()
}
func g() {
	defer func() {
		fmt.Println("g() defer")
	}()
	f()
	// 此后的代码不会被执行
	fmt.Println("after call f()")
}
func main() {
	g()
}
```

输出:

~~~
f() before panic
f() defer
g() defer
panic: 0

goroutine 1 [running]:
main.f()
	/tmpfs/gosandbox-28b34fa2_537076fc_b472688a_9fdc0e9b_cba83d95/prog.go:12 +0xc9
main.g()
	/tmpfs/gosandbox-28b34fa2_537076fc_b472688a_9fdc0e9b_cba83d95/prog.go:23 +0x2f
main.main()
	/tmpfs/gosandbox-28b34fa2_537076fc_b472688a_9fdc0e9b_cba83d95/prog.go:28 +0x18
~~~