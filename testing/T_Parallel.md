## func (t *T) Parallel()

参数列表

  <无>

返回值：

  <无>

功能说明：

调用Parallel()方法使当前测试与（且仅与）其他可并列运行的（同样调用了Parallel()方法的）测试并列运行于当前 CPU 组中。

代码实例：

	package testing_demo

	import (
		"testing"
	)

	func Test(t *testing.T) {
		t.Parallel()
	}

运行方法：

将上面的实例代码保存在文件“$GOPATH/src/testing_demo/t_parallel_test.go”中。用命令行进入文件所在目录并运行命令“go test -v”。
