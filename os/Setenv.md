## func Setenv(key, value string) error

参数列表

- key 环境变量名
- value 环境变量名的值

返回值：

- 返回error 返回错误信息对象

功能说明：

这个函数主要是设置一个环境变量

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fmt.Printf("Path of go is: %s\n", os.Getenv("go"))
        if err := os.Setenv("go", "/path/to/go"); err != nil {
            fmt.Printf("Error: %v", err)
        }
        fmt.Printf("Now Path of go is: %s\n", os.Getenv("go"))
    }

代码输出：

    //test in ArchLinux
    Path of go is:
    Now Path of go is: /path/to/go
