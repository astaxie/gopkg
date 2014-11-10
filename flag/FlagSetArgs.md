## func (f *FlagSet) Args() []string

参数列表

返回值
- []string 所有非flag参数的集合

功能说明
- 获取Flag集合f中所有非flag参数的集合

代码示例1
    
    package main
    
    import (
        "flag"
        "fmt"
    )
    
    var myFlagSet = flag.NewFlagSet("myflagset", flag.ExitOnError)
    var stringFlag = myFlagSet.String("abc", "default value", "help mesage")
    
    func main() {
        myFlagSet.Parse([]string{"-abc", "def", "ghi", "123"})
        args := myFlagSet.Args()
        for i := range args {
            fmt.Println(i, myFlagSet.Arg(i))
        }
    }

执行结果
    
    //  ./testflagsetarg
    0 ghi
    1 123
