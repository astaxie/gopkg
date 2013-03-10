package examples

import (
    "fmt"
    "encoding/base64"
)

func ExampleDecodeString1() {

    src := "dGhpcyBpcyBhIHRlc3Qgc3RyaW5nLg=="
    dst, _ := base64.StdEncoding.DecodeString(src)
    fmt.Println(string(dst) == "this is a test string.")

    dst, _ = base64.StdEncoding.DecodeString("")
    fmt.Println(string(dst) == "")

    dst, _ = base64.StdEncoding.DecodeString("Zg==")
    fmt.Println(string(dst) == "f")

    dst, _ = base64.StdEncoding.DecodeString("Zm8=")
    fmt.Println(string(dst) == "fo")

    dst, _ = base64.StdEncoding.DecodeString("Zm9v")
    fmt.Println(string(dst) == "foo")

    dst, _ = base64.StdEncoding.DecodeString("Zm9vYg==")
    fmt.Println(string(dst) == "foob")

    dst, _ = base64.StdEncoding.DecodeString("Zm9vYmE=")
    fmt.Println(string(dst) == "fooba")

    dst, _ = base64.StdEncoding.DecodeString("Zm9vYmFy")
    fmt.Println(string(dst) == "foobar")

    // Output:
    // true
    // true
    // true
    // true
    // true
    // true
    // true
    // true

}
