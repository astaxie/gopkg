## func (b *B) ResetTimer()

参数列表

  <无>

返回值：

  <无>

功能说明：

ResetTimer方法设置基准测试耗时为0。调用此方法不会影响到计时器的运行。

代码实例：

	import (
		"testing"
		"time"
	)

	func Benchmark(b *testing.B) {
		resetTag := false
		time.Sleep(time.Second)
		b.SetBytes(1000)
		if resetTag {
			b.ResetTimer()
		}
	}

运行方法：

将上面的实例代码保存在文件“$GOPATH/src/testing_demo/b_resettimer_test.go”中。用命令行进入文件所在目录并运行命令“go test -bench="." -v”。更改变量resetTag的值（true or false），再试试看。

备注：

按照运行方法反复运行代码实例，我们可以观察到：当resetTag为true时，此基准测试函数可以被运行多次，而当resetTag为false时，此基准测试函数往往只能获得一次运行机会。这是因为testing包中设置了这样一种限制：当基准测试函数单次运行时间超过指定值（默认为1秒，也可以根据命令行参数指定）时，只运行此基准测试函数一次。

因此，当我们调用time.Sleep()以强制延长函数运行时间之后，此基准测试函数就只会被运行一次了。而当我们在强制延时之后又调用了ResetTimer()时，相当于重置了此基准测试函数的运行时间，也即忽略了之前的强制延时操作。所以，在这种情况下，此基准测试函数自然就能运行多次了。

ResetTimer()方法的作用就在于此，将一些非关键的代码执行时间从基准测试函数运行时间中减去。这样也许可以获得更好的基准测试效果。
