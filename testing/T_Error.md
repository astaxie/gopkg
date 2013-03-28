## func (c *T) Error(args ...interface{})

参数列表

- args 要输出的、用于描述错误的内容

返回值：

  <无>

功能说明：

调用Error()相当于在调用了Log()之后再调用Fail()。

代码实例：

	import (
		"testing"
	)

	func Test(t *testing.T) {
		t.Error("Error:", "E")
	}

运行方法：

将上面的实例代码保存在文件“$GOPATH/src/testing_demo/t_error_test.go”中。用命令行进入文件所在目录并运行命令“go test -v”。
