## func (c *B) Fail()

参数列表

  <无>

返回值：

  <无>

功能说明：

调用Fail()使得当前测试函数失败，但扔会继续执行后面的代码。

代码实例：

    package testing_demo

	import (
		"testing"
	)

	func Benchmark(b *testing.B) {
		b.Fail()
	}

运行方法：

将上面的实例代码保存在文件“$GOPATH/src/testing_demo/b_fail_test.go”中。用命令行进入文件所在目录并运行命令“go test -bench="." -v”。
