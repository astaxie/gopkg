## func OpenFile(name string, flag int, perm FileMode) (file *File, err error)

参数列表

- name 文件名
- flag 文件打开方式
- perm 文件权限

返回值：

- file 文件指针
- error 返回错误信息对象

功能说明：

这个函数主要是打开一个文件，方式和权限通过参数控制

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fi, err := os.OpenFile("Hello.go", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0777)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            return
        }
        defer fi.Close()
        data := make([]byte, 100)
        fi.Read(data)
        fmt.Printf("The %s's content is: %s \n", fi.Name(), data)
        fi.WriteString("come on!!\n")
        fi.Seek(0, 0)
        fi.Read(data)
        fmt.Printf("Now the %s's content is: %s \n", fi.Name(), data)
    }

代码输出：

    //test in ArchLinux
    The Hello.go's content is: Hello World!

    Now the Hello.go's content is: Hello World!
    come on!!
