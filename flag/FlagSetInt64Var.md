## func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string)

参数列表
- p *int64 需要与flag参数值绑定的变量地址
- name string  flag名称
- value int 变量默认值
- usage string 提示信息

返回值

功能说明
- 将f中指定flag参数值绑定到int64变量

功能说明
- 为flag集合f增加爱一个带默认值和提示语句的int64类型flag，返回flag对应变量值的地址

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
    	int64Flag = myFlagSet.Int64("int64Flag", 100, defaultUsage)
    	int64Var  int64
    )
    
    func init() {
    	myFlagSet.Int64Var(&int64Var, "int64Var", 20, defaultUsage)
    }
    
    func main() {
    	args := []string{
    		"--int64Flag", "22",
    		"-int64Var", "-10",
    		"arg2",
    	}
    	myFlagSet.Parse(args)
    	fmt.Println("int64Flag", *int64Flag)
    	fmt.Println("int64Var", int64Var)
    }

代码输出
     
    //  ./testflagsetint64 
    int64Flag 22
    int64Var -10
