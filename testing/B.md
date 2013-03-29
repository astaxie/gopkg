## type B

结构：

	type B struct {
	    N int
	    // 包括被过滤或未被导出的字段
	}

类型说明：

B是一个类型，它会被传递给Benchmark函数，用于管理基准测试定时和指定迭代运行次数（N）。

代码实例：

    package testing_demo

	import (
		"testing"
	)

	func Benchmark(b *testing.B) {
		b.Logf("N: %d\n", b.N)
	}

运行方法：

将上面的实例代码保存在文件“$GOPATH/src/testing_demo/b_test.go”中。用命令行进入文件所在目录并运行命令“go test -bench="." -v”，观察输出，相信你会懂的。
