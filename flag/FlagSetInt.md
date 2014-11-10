##  func (f *FlagSet) Int(name string, value int, usage string) *int

参数列表
- name string   flag名称
- value int 变量默认值
- usage string 提示信息

返回值
- *int 返回一个int类型的flag值的地址

功能说明
- 为flag集合f增加一个带默认值和提示语句的int类型flag，返回对应值的地址

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
