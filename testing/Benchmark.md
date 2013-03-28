## func Benchmark(f func(b *B)) BenchmarkResult

参数列表

- f 要被测试的函数

返回值：

- BenchmarkResult实例值，基准测试结果。

功能说明：

Benchmark()用来测试单一函数。它可以被用于在不使用“go test”命令的情况下创建自定义的基准测试。

代码实例：

	package main

	import (
		"fmt"
		"testing"
	)

	func main() {
		br := testing.Benchmark(func(b *testing.B) { b.Log("B!") })
		fmt.Printf("Benchmark Result: %v\n", br)
	}
