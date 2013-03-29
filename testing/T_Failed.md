## func (c *B) Failed() bool

参数列表

  <无>

返回值：

- bool值，测试函数是否已失败。

功能说明：

判断测试函数是否已失败。

代码实例：

	package testing_demo

	import (
		"testing"
	)

	func Test(t *testing.T) {
		t.Logf("Failed: %v\n", t.Failed())
	}

运行方法：

将上面的实例代码保存在文件“$GOPATH/src/testing_demo/t_failed_test.go”中。用命令行进入文件所在目录并运行命令“go test -v”。
