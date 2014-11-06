## func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet

参数列表
- name string   flagset的名称
- errorHandling ErrorHandling flagset的错误处理方式，包括`ContinueOnError`:出错仍继续, `ExitOnError`:出错后退出程序,`PanicOnError`:出错后panic三种错误处理方式

返回值
- *FlagSet Flagset指针

功能说明
- 获取一个指定名称和错误处理方式的空FlagSet

代码示例
    
    package main
    
    import (
        "flag"
        "fmt"
    )
    
    func main() {
        var myFlagSet = flag.NewFlagSet("myFlagSet", flag.ExitOnError)
        fmt.Println(myFlagSet)
    }

执行结果
    
    //  ./testnewflagset 
    &{<nil> myFlagSet false map[] map[] [] 1 ?reflect.Value?}