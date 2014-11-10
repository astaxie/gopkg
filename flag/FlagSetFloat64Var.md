## func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string)

参数列表
- p *float64 需要与flag参数值绑定的变量地址
- name string   flag名称
- value float64 变量默认值
- usage string 提示信息

返回值

功能说明
- 将f中指定flag参数值绑定到float64变量

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
        
    //  ./testflagsetfloat64 
    float64Flag 2.718e+31
    float64Var -12.78
