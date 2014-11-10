## func (f *FlagSet) VisitAll(fn func(*Flag))

参数列表
- fn func(*Flag)  函数

返回值

功能说明
- 按照字典顺序遍历所有已定义的flag参数(不管是否已设置)，对每个参数调用fn方法

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
    	myFlagSet   = flag.NewFlagSet("newflagset", flag.ExitOnError)
    	uintFlag    = myFlagSet.Uint("uintFlag", 1000, defaultUsage)
    	testFlag    = myFlagSet.String("testFlag", "default value", defaultUsage)
    	float64Flag = myFlagSet.Float64("float64Flag", 0.9917, defaultUsage)
    	uintVar     uint
    )
    
    func print(f *flag.Flag) {
    	fmt.Println("	", f.Name, f.Value)
    }
    
    func main() {
    	args := []string{
    		"-uintFlag", "24",
    		"-float64Flag", "2718e28",
    		"arg1",
    	}
    	myFlagSet.Parse(args)
    	fmt.Println("Visit:")
    	myFlagSet.Visit(print)
    	fmt.Println("VisitAll:")
    	myFlagSet.VisitAll(print)
    }

运行结果
    
    //  ./testflagsetvisit 
    Visit:
    	 float64Flag 2.718e+31
    	 uintFlag 24
    VisitAll:
    	 float64Flag 2.718e+31
    	 testFlag default value
    	 uintFlag 24

