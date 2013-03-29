## type InternalTest

结构：

	type InternalTest struct {
	    Name string
	    F    func(*T)
	}

类型说明：

一个内部类型，只为了跨包使用而成为可导出的；它是“go test”命令实现的一部分。

代码实例：

	package main

	import (
		"fmt"
		"testing"
	)

	func main() {
		iTest := testing.InternalTest{Name: "DemoTest", F: func(t *testing.T) {}}
	    fmt.Printf("Test: %v\n", iTest)
	}
