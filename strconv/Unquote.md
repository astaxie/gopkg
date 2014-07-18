# func Unquote(s string) (t string, err error)

参数列表

- s     被单引号、双引号或反引号引用的字符串

返回值：

- t     将s转换为被引用前的字符串
-err

功能说明：

- 返回将被（单引号、双引号、反引号）引用的字符串的原字符串。（如果s是被单引号引用，则被引用的字符串必须是go中的字符，并返回该单字符的字符串。）

代码实例：

    package main
    
    import (
        "fmt"
        "strconv"
    )
    
    func main() {
        test := func(s string) {
            t, err := strconv.Unquote(s)
            if err != nil {
                fmt.Printf("Unquote(%#v) error： %v\n", s, err)
            } else {
                fmt.Printf("Unquote(%#v) = %v\n", s, t)
            }
        }
    
        s := `cafe\u0301`
        // If the string doesn't have quotes, it can't be unquoted.
        test(s) // invalid syntax
        test("`" + s + "`")
        test(`"` + s + `"`)
    
        test(`'\u00e9'`)
    
    }

代码输出：

    Unquote("cafe\\u0301") error： invalid syntax
    Unquote("`cafe\\u0301`") = cafe\u0301
    Unquote("\"cafe\\u0301\"") = café
    Unquote("'\\u00e9'") = é