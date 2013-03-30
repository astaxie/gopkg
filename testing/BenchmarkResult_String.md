## func (r BenchmarkResult) String() string

参数列表

  <无>

返回值：

- string值，本BenchmarkResult实例值的字符串形式。

功能说明：

获取单次迭代的耗时（单位：纳秒）。

代码实例：

	package testing_demo

	import (
		"testing"
		"time"
	)

	func Benchmark(b *testing.B) {
		br := testing.BenchmarkResult{N: 1, T: time.Second, Bytes: 1000}
		b.Logf("Benchmark    %d    %d ns/op    %d MB/s\n", br.N, br.T, br.Bytes)
		b.Logf("Benchmark String: %s\n", br.String())
	}


运行方法：

将上面的实例代码保存在文件“$GOPATH/src/testing_demo/benchmarkresult_string_test.go”中。用命令行进入文件所在目录并运行命令“go test -bench="." -v”。
