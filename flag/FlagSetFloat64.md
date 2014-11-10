## func (f *FlagSet) Float64(name string, value float64, usage string) *float64


参数列表

- name string flag名称
- value float64 默认值
- usage string 提示信息

返回值
- *float64 返回一个float64类型的flag值的地址

功能说明
- 为flag集合f增加一个带默认值和语句的float64类型的flag，返回flag对应值的地址

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
    	float64Flag = myFlagSet.Float64("float64Flag", 0.9917, defaultUsage)
    	float64Var  float64
    )
    
    func init() {
    	myFlagSet.Float64Var(&float64Var, "float64Var", 0.102, defaultUsage)
    }
    
    func main() {
    	args := []string{
    		"-float64Flag", "2718e28",
    		"-float64Var", "-12.78",
    		"arg2",
    	}
    	myFlagSet.Parse(args)
    	fmt.Println("float64Flag", *float64Flag)
    	fmt.Println("float64Var", float64Var)
    }

代码输出
        
./testflagsetfloat64 
float64Flag 2.718e+31
float64Var -12.78
