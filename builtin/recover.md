## func recover() interface{}

参数列表：

- 无

返回值：

- interface{}

功能说明：

recover 内建函数使程序可以管理恐慌中的 goroutine 。 recover 在推迟函数（其它函数总是返回 nil）中调用，通过执行 recover 停止恐慌过程序列并取回传至 panic 的参数值，恢复正常的执行。若 recover 不在推迟函数中被调用， 它将不会停止恐慌过程序列。在此情况下或该 goroutine 不在恐慌过程中，或提供给 panic 的实参为 nil ，recover 就会返回 nil。因此 recover 的返回值就报告了该 goroutine 是否进入恐慌过程序列。

代码实例：

```go
// Go Try catch 版 Hello World
package main

import (
	"fmt"
)

func Try(fun func(), catch func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catch(err)
		}
	}()
	fun()
}
func say(s string) {
	fmt.Println(s)
}
func main() {
	say("Hello")
	Try(
		func() {
			panic("World")
		},
		func(e interface{}) {
			fmt.Println("catch", e)
		})
	say("end")
}

```

输出:

~~~
Hello
catch World
end
~~~
