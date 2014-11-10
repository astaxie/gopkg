## func (f *FlagSet) Init(name string, errorHandling ErrorHandling)

参数列表
- name string 名称
- errorHandling ErrorHandling 错误处理方式，包括`ContinueOnError`:出错仍继续, `ExitOnError`:出错后退出程序,`PanicOnError`:出错后panic三种错误处理方式

返回值

功能说明
- 设置flag集合的名称和错误处理方式

代码示例
        
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var myFlagSet = flag.NewFlagSet("testFlagSet", flag.ExitOnError)
    
    func main() {
    	fmt.Println("Before Set:", myFlagSet)
    	myFlagSet.Init("myFlagSet", flag.ContinueOnError)
    	fmt.Println("After Set:", myFlagSet)
    }

代码输出
        
    //  ./testflagsetinit 
    Before Set: &{<nil> testFlagSet false map[] map[] [] 1 ?reflect.Value?}
    After Set: &{<nil> myFlagSet false map[] map[] [] 0 ?reflect.Value?}
