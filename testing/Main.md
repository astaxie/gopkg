## func Main(matchString func(pat, str string) (bool, error), tests []InternalTest, benchmarks []InternalBenchmark, examples []InternalExample)

参数列表

- matchString 测试名称匹配函数 
- tests 内部测试结构列表
- benchmarks 内部基准测试（性能测试）结构列表
- examples 内部示例测试结构列表

返回值：

  <无>

功能说明：

一个用于运行所有类型测试的内部函数，只为了跨包使用而成为可导出的；它是“go test”命令实现的一部分。

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
		iTest := testing.InternalTest{Name: "DemoTest", F: func(t *testing.T) {}}
		iTests := make([]testing.InternalTest, 1)
		iTests[0] = iTest
		iBenchmark := testing.InternalBenchmark{Name: "DemoBenchMark", F: func(b *testing.B) {}}
		iBenchmarks := make([]testing.InternalBenchmark, 1)
		iBenchmarks[0] = iBenchmark
		iExample := testing.InternalExample{Name: "Example", F: func() {}, Output: ""}
		iExamples := make([]testing.InternalExample, 1)
		iExamples[0] = iExample
		testing.Main(matchString, iTests, iBenchmarks, iExamples)
	}
