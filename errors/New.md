# func New(text string) error

参数列表

- text 错误字符串

返回值：

- 返回error

功能说明：

New returns an error that formats as the given text.
按给定的文本返回一个新的 error.

代码实例：

        package main

        import  "fmt"
        import  "errors"

        func main(){
                fmt.Println(errors.New("Err"))
        }

