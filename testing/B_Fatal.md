## func (c *B) Fatal(args ...interface{})

参数列表

- args 要输出的、用于描述错误的内容

返回值：

  <无>

功能说明：

调用Error()相当于在调用了Log()之后再调用FailNow()。

代码实例：

	package testing_demo

	import (
		"testing"
	)

	func Benchmark(b *testing.B) {
		b.Fatal("Fatal:", "F")
	}

运行方法：

将上面的实例代码保存在文件“$GOPATH/src/testing_demo/b_fatal_test.go”中。用命令行进入文件所在目录并运行命令“go test -bench="." -v”。
