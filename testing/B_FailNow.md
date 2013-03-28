## func (c *B) FailNow()

参数列表

  <无>

返回值：

  <无>

功能说明：

调用Fail()使得当前测试函数失败且停止执行，转而执行后面的测试函数或基准测试函数。

代码实例：

    package testing_demo

	import (
		"testing"
	)

	func Benchmark(b *testing.B) {
		b.FailNow()
	}

运行方法：

将上面的实例代码保存在文件“$GOPATH/src/testing_demo/b_failnow_test.go”中。用命令行进入文件所在目录并运行命令“go test -bench="." -v”。
