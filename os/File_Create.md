## func Create(name string) (file *File, err error)

参数列表

- name 文件名

返回值：

- file 文件指针
- error 返回错误信息对象

功能说明：

这个函数主要是创建一个文件，此方法会覆盖已存在的文件

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fi, err := os.Create("Hello.go")
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            return
        }
        defer fi.Close()
        fmt.Printf("The %s has been created!\n", fi.Name())
    }

代码输出：

    //test in ArchLinux
    The Hello.go has been created!
