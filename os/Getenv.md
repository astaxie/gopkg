## func Getenv(key string) string

参数列表

- key 环境变量的key

返回值：

- string 返回环境变量的key对应的值

功能说明：

这个函数主要是根据key来获取相当环境变量的值

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fmt.Printf("%s\n", os.Getenv("LANG"))
    }

代码输出：

    //test in ArchLinux
    zh_CN.UTF-8
