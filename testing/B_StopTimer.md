## func (b *B) StopTimer()

参数列表

  <无>

返回值：

  <无>

功能说明：

StartTimer方法停止对一个测试的计时。当你有不想纳入的复杂初始化执行时间时，可以调用此方法以暂停计时器。

代码实例：

	import (
		"testing"
		"time"
	)

	func Benchmark(b *testing.B) {
		customTimerTag := false
		if customTimerTag {
			b.StopTimer()
		}
		time.Sleep(time.Second)
		if customTimerTag {
			b.StartTimer()
		}
	}

运行方法：

将上面的实例代码保存在文件“$GOPATH/src/testing_demo/b_stoptimer_test.go”中。用命令行进入文件所在目录并运行命令“go test -bench="." -v”。更改变量customTimerTag的值（true or false），再试试看。


备注：

实例中的代码（实际上是与方法StartTimer()配合使用的）与调用ResetTimer()方法有异曲同工之妙，具体的用途和作用可以参看本文档中[对方法ResetTimer()的说明](B_ResetTimer.md)。
