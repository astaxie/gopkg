## func Glob(pattern string) (matches []string, err error)

参数列表

- pattern 匹配的模式

返回值：

- matches 所有匹配的文件 
- 返回bool error

功能说明：

返回目录下所有匹配的文件


示例:

    package main

    import (
        "fmt"
        "path/filepath"
    )


    func main(){
        //返回目录下所有的go文件
        matches, _ := filepath.Glob("*.go")

        //找出/home/ 目录下的所有的log文件 
        matches, _ = filepath.Glob("/home/*.log")
    }
