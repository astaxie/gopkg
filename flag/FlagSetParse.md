## func (f *FlagSet) Parse(arguments []string) error

参数列表
－ arguments []string 要解析的字符串列表

返回值

功能说明
- 从argumets中依次解析flag.在定义flag之后，必须调用用本函数才能成功获取arguments中设置的flag的值 

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
    	testFlag  = myFlagSet.String("testFlag", "default value", defaultUsage)
    	uintVar   uint
    )
    
    func init() {
    	myFlagSet.UintVar(&uintVar, "uintVar", 30, defaultUsage)
    }
    
    func main() {
    	args := []string{
    		"-uintFlag", "24",
    		"-testFlag", "test",
    		"arg1",
    	}
    	myFlagSet.Parse(args)
    	if myFlagSet.Parsed() {
    		fmt.Println("uintFlag:", *uintFlag)
    		fmt.Println("testFlag:", *testFlag)
    	}
    }

代码输出
        
    //  ./testflagsetparse                          
    uintFlag: 24
    testFlag: test


