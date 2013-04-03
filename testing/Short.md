## func Short() bool

参数列表

<无>

返回值：

- bool值，指定的flag是否被设置。

功能说明：

判断flag "-test.short" 是否被设置。
这一flag被设置后可以使测试跑得更快，但其具体功能由测试编写者自己实现。testing包仅仅作为其承载者。在安装脚本all.bash中设置了这一flag以使安装更有效率。但是在执行“go test”命令时，这一标记在默认情况下是不会被设置的，这使得当前包的完整测试会被运行。 

代码实例：

	package testing_demo

	import (
		"testing"
	)

	func TestShort(t *testing.T) {
		t.Logf("The short flag: %v\n", testing.Short())
	}

运行方法：

将上面的实例代码保存在文件“$GOPATH/src/testing_demo/short_test.go”中。用命令行进入文件所在目录并运行命令“go test -test.short -v”。删除命令中的字符串“-test.short”再看看效果。
