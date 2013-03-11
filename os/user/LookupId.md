## func LookupId(uid string) (*User, error)

参数列表

-uid 要查询的用户的uid

返回值：

- 返回User 查询到用户对象
- 返回error 返回错误信息对象

功能说明：

这个函数主要是用来根据uid来查询用户对象

代码实例：

    package main

    import (
        "fmt"
        "os/user"
    )

    func main() {
        findedUser, err := user.LookupId("501")
        if err != nil {
            fmt.Println(err)
        }
        //我所在环境为Mac os X 10.8.2
        fmt.Println("用户id:", findedUser.Uid)      //用户id: 501
        fmt.Println("用户所在组id:", findedUser.Gid)   //    用户所在组id: 20
        fmt.Println("用户名:", findedUser.Username)  //用户名: lsdev
        fmt.Println("用户全名:", findedUser.Name)     //用户全名: 吴 佳煌
        fmt.Println("用户家目录:", findedUser.HomeDir) //用户家目录: /Users/lsdev

    }   

