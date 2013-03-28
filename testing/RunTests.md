## func RunTests(matchString func(pat, str string) (bool, error), tests []InternalTest) (ok bool)

参数列表

- matchString 测试名称匹配函数 
- tests 内部测试结构列表
- benchmarks 内部基准测试（性能测试）结构列表
- examples 内部示例测试结构列表

返回值：

- ok bool值，是否通过测试。

功能说明：

运行测试。

代码实例：

	import (
		"fmt"
		"regexp"
		"testing"
	)

	func matchString(pat, str string) (bool, error) {
		return regexp.MatchString(pat, str)
	}

	func main() {
		iExample := testing.InternalExample{Name: "Example", F: func() {}, Output: ""}
		iExamples := make([]testing.InternalExample, 1)
		iExamples[0] = iExample
		done := testing.RunExamples(matchString, iExamples)
		fmt.Printf("Result: %v\n", done)
	}
