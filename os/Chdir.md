## func Chdir(dir string) error

参数列表

- dir 目录路径

返回值：

- 返回error 返回错误信息对象

功能说明：

这个函数主要是改变当前工作目录

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        pwd, err := os.Getwd()
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        fmt.Printf("The current directory is: %s\n", pwd)

        if err = os.Chdir("go-study"); err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }

        pwd, err = os.Getwd()
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        fmt.Printf("Now the current directory is: %s\n", pwd)
    }

代码输出：

    //test in ArchLinux
    //go-study为当前文件所有目录下的一个子目录
    The current directory is: /home/miraclesu/Programs/go
    Now the current directory is: /home/miraclesu/Programs/go/go-study
