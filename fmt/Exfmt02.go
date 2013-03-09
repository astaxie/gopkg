// Exfmt02.go
// 测试和说明：
// %t 值的true和false
// %T 值的类型在Go语言中的表示
// %% 打印一个 %
// %q 给打印的字串自动加引号
package main

import "fmt"

func main() {
        var Yes bool //bool 变量自动初始化为 false
        fmt.Printf("No01:%t\n", Yes)
        fmt.Printf("No02:%T\n", Yes)
        fmt.Printf("No03:%%\n")
        fmt.Printf("No04:%q\n","自动加引号")
}
// 输出：
//No01:false
//No02:bool
//No03:%
//No04:"自动加引号"
