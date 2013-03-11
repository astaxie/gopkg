## func Contains(s, substr string) bool

参数列表



返回值：

- 返回User 当前用户对象
- 返回error 返回错误信息对象

功能说明：

这个函数主要是用来获取系统当前的用户对象

代码实例：

    package main

    import (
        "fmt"
        "os/user"
    )

    func main() {
        user, err := user.Current()
        if err != nil {
            fmt.Println(err)
        }
        //我所在环境为Mac os X 10.8.2
        fmt.Println("用户id:", user.Uid)      //用户id: 501
        fmt.Println("用户所在组id:", user.Gid)   //    用户所在组id: 20
        fmt.Println("用户名:", user.Username)  //用户名: lsdev
        fmt.Println("用户全名:", user.Name)     //用户全名: 吴 佳煌
        fmt.Println("用户家目录:", user.HomeDir) //用户家目录: /Users/lsdev

    }

