## func Visit(fn func(*Flag))

参数列表
- fn func(*Flag)  函数

返回值

功能说明
- 按照字典顺序遍历所有命令行已设置的flag参数(只定义了但未在命令行被设置的flag不会被遍历)，对每个参数调用fn方法

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
    	flag.Visit(print)
    }

运行结果
    
    //  ./testvisit -string "teststring" -int -90 
    int -90
    string teststring
