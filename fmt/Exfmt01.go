// Exfmt01.go
// 测试和说明：
// %v  基本格式的值
// %+v 当输出结构体时，扩展标志添加成员的名字。
// %#v 值的Go语法表示。
package main

import "fmt"

var timeZone = map[string]int{
        "UTC": 0 * 60 * 60,
        "EST": -5 * 60 * 60,
        "CST": -6 * 60 * 60,
        "MST": -7 * 60 * 60,
        "PST": -8 * 60 * 60,
}

type T struct {
        a int
        b float32
        c string
}

func main() {
        t := &T{7, -2.35, "abc\tdef"}
        fmt.Printf("No01:%v\n", t)
        fmt.Printf("No02:%+v\n", t)
        fmt.Printf("No03:%#v\n", t)
        fmt.Printf("No04:%#v\n", timeZone)
}
// 输出：
//No01:&{7 -2.35 abc    def}
//No02:&{a:7 b:-2.35 c:abc      def}
//No03:&main.T{a:7, b:-2.35, c:"abc\tdef"}
//No04:map[string]int{"MST":-25200, "UTC":0, "CST":-21600, "PST":-28800, "EST":-18000}
