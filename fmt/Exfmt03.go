// Exfmt03.go
// 测试和说明：
// 整数部分
// %b 二进制表示
// %c 数值对应的Unicode编码字符
// %d 十进制表示
// %o 八进制表示
// %x 十六进制表示，使用a-f
// %X 十六进制表示，使用A-F
// %U Unicode格式： U+1234，等价于"U+%04X"
// %q 单引号
package main

import "fmt"

func main() {
        fmt.Printf("No01:%b\n", 888)
        fmt.Printf("No02:%c\n", '太')
        fmt.Printf("No03:%d\n", 888)
        fmt.Printf("No04:%o\n", 888)
        fmt.Printf("No05:%x\n", 888)
        fmt.Printf("No06:%X\n", 888)
        fmt.Printf("No07:%U\n", 888)
        fmt.Printf("No08:%q\n", 88)
}
// 输出：
//No01:1101111000
//No02:太
//No03:888
//No04:1570
//No05:378
//No06:378
//No07:U+0378
//No08:'X'
