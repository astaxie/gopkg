## type T

结构：

	type T struct {
	    // 包括被过滤或未被导出的字段
	}

类型说明：

T是一个类型，它会被传递给Test函数，用于管理测试状态并支持格式化了的测试日志。日志会在测试执行及标准错误被抛出的时候累积，并在测试完成时一并输出。

代码实例：

	package testing_demo

	import (
		"testing"
	)

	func Test(t *testing.T) {
		t.Log("Testing...")
	}

运行方法：

将上面的实例代码保存在文件“$GOPATH/src/testing_demo/t_test.go”中。用命令行进入文件所在目录并运行命令“go test -v”。
