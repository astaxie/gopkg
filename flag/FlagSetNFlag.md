## func (f *FlagSet) NFlag() int

参数列表

返回值
- int 返回解析成功的flag个数

功能说明
- 获取解析成功的flag个数

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
    		"arg1",
    	}
    	fmt.Println("before:", myFlagSet.NFlag())
    	myFlagSet.Parse(args)
    	fmt.Println("after:", myFlagSet.NFlag())
    }

代码输出
        
    //  ./testflagsetnflag                          
    before: 0
    after: 1
