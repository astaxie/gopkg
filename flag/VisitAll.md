## func VisitAll(fn func(*Flag))

参数列表
- fn func(*Flag)  函数

返回值

功能说明
- 按照字典顺序遍历所有已定义的flag参数(不管是否在命令行中设置)，对每个参数调用fn方法

代码示例
    
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var stringFlag = flag.String("string", "string", "help message for string")
    var intFlag = flag.Int("int", 100, "help message for int")
    var boolFlag = flag.Bool("bool", false, "help message for bool")
    var testFlag = flag.String("test", "test", "help message for test")
    
    func print(f *flag.Flag) {
    	fmt.Println(f.Name, f.Value)
    }
    
    func main() {
    	flag.Parse()
    	flag.VisitAll(print)
    }

运行结果
    
    //  ./testvisitall -string "teststring" -int 999
    bool false
    int 999
    string teststring
    test test

