##  func (f *FlagSet) IntVar(p *int, name string, value int, usage string)

参数列表
- p *int 需要与flag参数值绑定的变量地址
- name string  flag名称
- value int 变量默认值
- usage string 提示信息

返回值

功能说明
- 将f中指定flag参数值绑定到int64变量

代码示例
    
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    const (
    	defaultUsage = "help message"
    )
    
    var (
    	myFlagSet = flag.NewFlagSet("newflagset", flag.ExitOnError)
    	intFlag   = myFlagSet.Int("intFlag", 100, defaultUsage)
    	intVar    int
    )
    
    func init() {
    	myFlagSet.IntVar(&intVar, "intVar", 20, defaultUsage)
    }
    
    func main() {
    	args := []string{
    		"--intFlag", "22",
    		"-intVar", "-10",
    		"arg2",
    	}
    	myFlagSet.Parse(args)
    	fmt.Println("intFlag", *intFlag)
    	fmt.Println("intVar", intVar)
    }

代码输出
        
    //  ./testflagsetint
    intFlag 22
    intVar -10
