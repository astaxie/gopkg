## func (f *FlagSet) NArg() int

参数列表

返回值
- int 返回解析后剩余的参数的数量

功能说明
- 获取输入参数解析后剩余的参数的数量，即f.Args()返回的值的元素个数

示例
        
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
    	uintFlag  = myFlagSet.Uint("uintFlag", 1000, defaultUsage)
    	uintVar   uint
    )
    
    func init() {
    	myFlagSet.UintVar(&uintVar, "uintVar", 30, defaultUsage)
    }

    func main() {
    	args := []string{
    		"-uintFlag", "24",
    		"--uintVar", "25",
    		"arg1",
    		"arg2",
    		"arg3",
    	}
    	myFlagSet.Parse(args)
    	length := myFlagSet.NArg()
    	for i := 0; i < length; i++ {
    		fmt.Println(i, myFlagSet.Arg(i))
    	}
    }

代码输出
        
    //  ./testflagsetnarg 
    0 arg1
    1 arg2
    2 arg3
