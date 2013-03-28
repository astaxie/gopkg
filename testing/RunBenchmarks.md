## func RunBenchmarks(matchString func(pat, str string) (bool, error), benchmarks []InternalBenchmark)

参数列表

- matchString 测试名称匹配函数 
- benchmarks 内部基准测试（性能测试）结构列表

返回值：

  <无>

功能说明：

一个用于运行基准测试的内部函数，只为了跨包使用而成为可导出的；它是“go test”命令实现的一部分。

代码实例：

	package main

	import (
		"regexp"
		"testing"
	)

	func matchString(pat, str string) (bool, error) {
		return regexp.MatchString(pat, str)
	}

	func main() {
		iBenchMark := testing.InternalBenchmark{Name: "DemoBenchMark", F: func(b *testing.B) {}}
		iBenchMarks := make([]testing.InternalBenchmark, 1)
		iBenchMarks[0] = iBenchMark
		testing.RunBenchmarks(matchString, iBenchMarks)
	}
