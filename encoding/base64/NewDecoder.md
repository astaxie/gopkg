## func NewDecoder(enc *Encoding, r io.Reader) io.Reader

参数列表：

- enc Encoding 结构指针
- w io.Reader 接口

返回值：

- io.Reader 接口

功能说明：

返回一个 io.Reader 接口，用于流式解码

代码实例：

    package main

    import (
        "fmt"
        "strings"
        "encoding/base64"
    )

    // 从 string 取得输入
    func Example1() {

        fmt.Println("---=== Example1 ===---")

        src := "dGhpcyBpcyBhIHRlc3Qgc3RyaW5nLg=="
        reader := strings.NewReader(src)

        dst := ""

        decoder := base64.NewDecoder(base64.StdEncoding, reader)
        // 使用一个很小的输出buffer，测试流式解码
        buf := make([]byte, 2)
        for {
            n, err := decoder.Read(buf)
            if err != nil {
                break
            }
            if n == 0 {
                break
            }
            dst += string(buf[0:n])
        }
        fmt.Print(dst) // this is a test string.
        fmt.Print("\n")
    }

    func main() {

        Example1()

    }
