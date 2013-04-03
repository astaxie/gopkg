package examples

import (
    "fmt"
    "os"
    "encoding/base32"
)

func ExampleNewEncoder1() {

    // 输出到标准输出
    encoder := base32.NewEncoder(base32.StdEncoding, os.Stdout)
    encoder.Write([]byte("this is a test string."))
    // 必须调用Close，刷新输出
    encoder.Close()

    // Output:
    // ORUGS4ZANFZSAYJAORSXG5BAON2HE2LOM4XA====

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

    encoder := base32.NewEncoder(base32.StdEncoding, buf)
    encoder.Write([]byte("this is a test string."))
    encoder.Close()

    fmt.Print(buf.S == "ORUGS4ZANFZSAYJAORSXG5BAON2HE2LOM4XA====")

    // Output:
    // true

}
