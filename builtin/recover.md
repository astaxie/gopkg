## func recover() interface{}

参数列表：

- 无

返回值：

- interface{}

功能说明：

recover 内建函数使程序可以管理恐慌中的 goroutine 。 recover 在推迟函数（其它函数总是返回 nil）中调用，通过执行 recover 停止恐慌过程序列并取回传至 panic 的参数值，恢复正常的执行。若 recover 不在推迟函数中被调用， 它将不会停止恐慌过程序列。在此情况下或该 goroutine 不在恐慌过程中，或提供给 panic 的实参为 nil ，recover 就会返回 nil。因此 recover 的返回值就报告了该 goroutine 是否进入恐慌过程序列。

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

func g(label int) {
	// 模拟 try catch 并继续执行
	catch := func() {
		if e := recover(); e != nil {
			fmt.Println("panicing has ", e)
			defer g(1)
		} else {
			fmt.Println("no panicing")
		}
		return
	}
	defer catch()
	if label == 1 {
		goto CATCH
	}
	f()
	if label == 0 {
		goto NEXT
	}
CATCH:
	fmt.Println("catch call f()")
NEXT:
	fmt.Println("end")
}
func main() {
	g(0)
}
```

输出:

~~~
f() before panic
f() defer
panicing has  0
catch call f()
end
no panicing
~~~