## func (c *B) Log(args ...interface{})

参数列表

- args 要输出的、用于描述错误的内容

返回值：

  <无>

功能说明：

Log方法用默认的格式来格式化给定的参数，类似于Println()，并在错误日志中记录文本。

代码实例：

	package testing_demo

	import (
		"testing"
	)

	func Benchmark(b *testing.B) {
		b.Log("Log:", "L")
	}

运行方法：

将上面的实例代码保存在文件“$GOPATH/src/testing_demo/b_log_test.go”中。用命令行进入文件所在目录并运行命令“go test -bench="." -v”。
