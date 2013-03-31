package examples

import (
    "fmt"
    "encoding/base32"
)

func ExampleDecodeString1() {

    src := "ORUGS4ZANFZSAYJAORSXG5BAON2HE2LOM4XA===="
    dst, _ := base32.StdEncoding.DecodeString(src)
    fmt.Println(string(dst) == "this is a test string.")

    dst, _ = base32.StdEncoding.DecodeString("")
    fmt.Println(string(dst) == "")

    dst, _ = base32.StdEncoding.DecodeString("MY======")
    fmt.Println(string(dst) == "f")

    dst, _ = base32.StdEncoding.DecodeString("MZXQ====")
    fmt.Println(string(dst) == "fo")

    dst, _ = base32.StdEncoding.DecodeString("MZXW6===")
    fmt.Println(string(dst) == "foo")

    dst, _ = base32.StdEncoding.DecodeString("MZXW6YQ=")
    fmt.Println(string(dst) == "foob")

    dst, _ = base32.StdEncoding.DecodeString("MZXW6YTB")
    fmt.Println(string(dst) == "fooba")

    dst, _ = base32.StdEncoding.DecodeString("MZXW6YTBOI======")
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
