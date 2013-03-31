## type BenchmarkResult

结构：

	type BenchmarkResult struct {
	    N     int           // 迭代的次数
	    T     time.Duration // 总耗时
	    Bytes int64         // 单次迭代处理的字节数
	}

类型说明：

基准测试的结果。

代码实例：

	package testing_demo

	import (
		"testing"
		"time"
	)

	func Benchmark(b *testing.B) {
		br := testing.BenchmarkResult{N: 1, T: time.Second, Bytes: 1000}
		b.Logf("Benchmark    %d    %d ns/op    %d MB/s\n", br.N, br.T, br.Bytes)
	}

运行方法：

将上面的实例代码保存在文件“$GOPATH/src/testing_demo/benchmarkresult_test.go”中。用命令行进入文件所在目录并运行命令“go test -bench="." -v”，观察输出，相信你会懂的。

备注：

运行命令“go test -bench="." <...>”之后所看到的输出结果就是根据结构BenchmarkResult实例来填充的。