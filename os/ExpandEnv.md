## func ExpandEnv(s string) string

参数列表

- s 需要替换的字符串

返回值：

- string 替换后的字符串

功能说明：

这个函数主要是把一个字符串里带$var或${var}这样的字符串替换成当前环境变量中的内容，如果没有对应的，则替换成空

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        s := "The language path is: $LANG"
        fmt.Printf("%s\n", os.ExpandEnv(s))
    }

代码输出：

    //test in ArchLinux
    The language path is: zh_CN.UTF-8
