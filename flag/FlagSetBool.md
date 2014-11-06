## func (f *FlagSet) Bool(name string, value bool, usage string) *bool

参数列表
- name string flag名称
- value bool 默认值
- usage string 提示信息

返回值

功能说明
将f中指定flag的值绑定到一个bool变量

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
    	boolFlag  = myFlagSet.Bool("boolFlag", false, defaultUsage)
    	boolVar   bool
    )
    
    func init() {
    	myFlagSet.BoolVar(&boolVar, "boolVar", false, defaultUsage)
    }
    
    func main() {
    	args := []string{
    		"-boolFlag",
    		"-boolVar=true",
    		"arg2",
    	}
    	myFlagSet.Parse(args)
    	fmt.Println("boolFlag", *boolFlag)
    	fmt.Println("boolVar", boolVar)
    }

执行结果
    
    //  ./testflagsetbool                         
    boolFlag true
    boolVar true

