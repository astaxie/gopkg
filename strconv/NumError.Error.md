# func (e *NumError) Error() string

返回值：

- 返回string     描述错误的字符串

功能说明：

- 返回描述NumError类型错误的字符串

代码实例：

    package main
    
    import (
        "fmt"
        "strconv"
    )
    
    func main() {
        err := errors.New("too large")
        er := strconv.NumError{Func: "anyfunc", Num: "1e100",Err:err}
        fmt.Println(er.Error())
    }

代码输出：

    strconv.anyfunc: parsing "1e100": too large