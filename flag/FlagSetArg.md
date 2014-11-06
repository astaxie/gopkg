## func (f *FlagSet) Arg(i int) string

参数列表
- i int 非flag参数集合的索引

返回值
- string 第i＋1个非flag参数的值

功能说明
- 获取Flag集合f中第i＋1个非flag参数的值

代码示例1
    
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var myFlagSet = flag.NewFlagSet("myflagset", flag.ExitOnError)
    var stringFlag = myFlagSet.String("abc", "default value", "help mesage")
    
    func main() {
    	myFlagSet.Parse([]string{"-abc", "def", "ghi", "123"})
    	args := myFlagSet.Args()
    	for i := range args {
    		fmt.Println(i, myFlagSet.Arg(i))
    	}
    }

执行结果
    
    //  ./testflagsetarg
    0 ghi
    1 123
