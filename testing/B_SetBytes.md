## func (b *B) SetBytes(n int64)

参数列表

- n 处理的字节数

返回值：

  <无>

功能说明：

SetBytes方法记录在一个单一操作中处理的字节数。如果此方法被调用，则基准测试结果中会包含于与 ns/op 和 MB/s 有关的内容。

代码实例：

	package testing_demo

	import (
		"testing"
		"time"
	)

	func Benchmark(b *testing.B) {
		time.Sleep(time.Millisecond)
		b.SetBytes(1000)
	}

运行方法：

将上面的实例代码保存在文件“$GOPATH/src/testing_demo/b_setbytes_test.go”中。用命令行进入文件所在目录并运行命令“go test -bench="." -v”。注释掉“b.SetBytes(1000)”这句代码，再试试看。

备注：

当在基准测试函数中调用SetBytes方法之后，测试结果中的 MB/S 前的数字表明一秒钟处理的字节数。调用SetBytes方法时所传入的参数完全是测试者自己给定的。测试者可以依此观察被测试对象（比如用于IO的操作）的性能。
