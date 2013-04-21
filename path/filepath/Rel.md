## func Rel(basepath, targpath string) (string, error)

参数列表:

- basepath: 基准路径
- targpath: 目标路径

返回值列表:

- 返回以basepath为基准的相对路径。 
- error

功能说明:
返回以basepath为基准的相对路径。Join(basepath, Rel(basepath, targpath)) 等于targpath。

示例:

    package main

    import (
        "fmt"
        "path/filepath"
    )

    func main(){
        // 返回 xuchdong/rel_demo.go
        path, _ := filepath.Rel("/home", "/home/xuchdong/rel_demo.go")
        fmt.Printf("%s\n", path)
    }
