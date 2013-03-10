package examples

import (
    "fmt"
    "os"
    "encoding/base64"
)

func Example1() {

    // 输出到标准输出
    encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
    encoder.Write([]byte("this is a test string."))
    // 必须调用Close，刷新输出
    encoder.Close()

    // Output:
    // dGhpcyBpcyBhIHRlc3Qgc3RyaW5nLg==

}

type StringWriter struct {
    S string
}
func (this *StringWriter) Write(p []byte) (n int, err error) {
    this.S += string(p);
    return len(p), nil
}

// 输出到自定义 Writer
func ExampleNewEncoder2() {

    buf := &StringWriter{}

    encoder := base64.NewEncoder(base64.StdEncoding, buf)
    encoder.Write([]byte("this is a test string."))
    encoder.Close()

    fmt.Print(buf.S == "dGhpcyBpcyBhIHRlc3Qgc3RyaW5nLg==")

    // Output:
    // true

}
