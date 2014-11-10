## func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string)

参数列表
- p *uint64 需要与flag参数值绑定的变量地址
- name string  flag名称
- value uint64 变量默认值
- usage string 提示信息

返回值

功能说明
- 将f中指定flag参数值绑定到uint64变量

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
         myFlagSet  = flag.NewFlagSet("newflagset", flag.ExitOnError)
         uint64Flag = myFlagSet.Uint64("uint64Flag", 3310000, defaultUsage)
         uint64Var  uint64
    )
   
    func init() {
         myFlagSet.Uint64Var(&uint64Var, "uint64Var", 719200, defaultUsage)
    }
   
    func main() {
         args := []string{
              "--uint64Var", "25",
              "-uint64Flag", "1231",
              "arg2",
         }
         myFlagSet.Parse(args)
         fmt.Println("uint64Flag", *uint64Flag)
         fmt.Println("uint64Var", uint64Var)
    }


代码输出
           
     //  ./testflagsetuint64
     uint64Flag 1231
     uint64Var 25