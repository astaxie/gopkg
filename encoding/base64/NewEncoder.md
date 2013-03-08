## func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser

参数列表

- enc Encoding 结构指针
- w io.Writer 接口

返回值：

- io.WriteCloser 接口

功能说明：

返回一个 io.WriteCloser 接口，用于流式编码

代码实例：

    package main

    import (
        "fmt"
        "os"
        "encoding/base64"
    )

    func Example1() {

        fmt.Println("---=== Example1 ===---")

        // 输出到标准输出
        encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
        encoder.Write([]byte("this is a test string."))
        // 必须调用Close，刷新输出
        encoder.Close() // dGhpcyBpcyBhIHRlc3Qgc3RyaW5nLg==

        fmt.Print("\n")
    }

    type StringWriter struct {
        S string
    }
    func (this *StringWriter) Write(p []byte) (n int, err error) {
        this.S += string(p);
        return len(p), nil
    }

    // 输出到自定义 Writer
    func Example2() {

        fmt.Println("---=== Example2 ===---")

        buf := &StringWriter{}

        // 输出到标准输出
        encoder := base64.NewEncoder(base64.StdEncoding, buf)
        encoder.Write([]byte("this is a test string."))
        encoder.Close()

        fmt.Print(buf.S == "dGhpcyBpcyBhIHRlc3Qgc3RyaW5nLg==") // true
        fmt.Print("\n")
    }

    func main() {

        Example1()
        Example2()

    }

