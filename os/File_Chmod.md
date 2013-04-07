## func (f *File) Chmod(mode FileMode) error

参数列表

- mode 文件修改的目标模式

返回值：

- 返回error 返回错误信息对象

功能说明：

这个函数主要是修改文件的模式

代码实例：

    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fi, err := os.Stat("Hello.go")
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        fmt.Printf("Hello.go's mode is: %s(%o)\n", fi.Mode(), fi.Mode()&0777)

        f, _ := os.Open("Hello.go")
        err = f.Chmod(0777)
        if err != nil {
            fmt.Printf("Error: %s\n", err)
        }
        f.Close()

        fi, err = os.Stat("Hello.go")
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }
        fmt.Printf("Now Hello.go's mode is: %s(%o)\n", fi.Mode(), fi.Mode()&0777)
    }

代码输出：

    //test in ArchLinux
    Hello.go's mode is: -rw-r--r--(644)
    Now Hello.go's mode is: -rwxrwxrwx(777)
