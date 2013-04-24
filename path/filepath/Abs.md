## func Abs(path string) (string, error)

参数列表:

- path: 

返回值列表:

- 返回所给路径的绝对路径string 
- error

功能说明：

返回所给的路径的的绝对路径。

示例:

    package main

    import (
        "fmt"
        "path/filepath"
    )

    func main(){
        // 当前路径为/home, 如下返回的path将会是/home/abs_demo.go
        path, _ := filepath.Abs("abs_demo.go")
        fmt.Printf("%v\n", path)
    }


