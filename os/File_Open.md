## func Open(name string) (file *File, err error)

参数列表

- name 文件名

返回值：

- file 文件指针
- error 返回错误信息对象

功能说明：

这个函数主要是以只读的模式打开一个文件

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fi, err := os.Open("Hello.go")
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            return
        }
        defer fi.Close()
        data := make([]byte, 100)
        fi.Read(data)
        fmt.Printf("The %s's content is: %s \n", fi.Name(), data)
    }

代码输出：

    //test in ArchLinux
    The Hello.go's content is: Hello World!
