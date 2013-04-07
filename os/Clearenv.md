## func Clearenv()

参数列表

- 无

返回值：

- 无

功能说明：

这个函数主要是清空当前环境变量

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fmt.Printf("The path is: %+v\n", os.Environ())
        os.Clearenv()
        fmt.Printf("Now the path is: %+v\n", os.Environ())
    }

代码输出：

    //test in ArchLinux
    //只放了一部分结果
    The path is: [USER=miraclesu LG_ALL=zh_CN.UTF-8]
    Now the path is: []
