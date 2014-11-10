## func (f *FlagSet) Lookup(name string) *Flag

参数列表
- name string  flag名称

返回值
- *Flag Flag指针

功能说明
- 获取flag集合f中名称为name值的flag指针，如果对应的flag不存在，返回nil

代码示例
    
    package main
    
    import (
    	"flag"
    	"fmt"
    )
    
    var myFlagSet = flag.NewFlagSet("testFlagSet", flag.ExitOnError)
    var testFlag = myFlagSet.String("test", "default value", "help message.")
    
    func print(f *flag.Flag) {
    	if f != nil {
    		fmt.Println(f.Value)
    	} else {
    		fmt.Println(nil)
    	}
    }
    
    func main() {
    	fmt.Print("test:")
    	print(myFlagSet.Lookup("test"))
    	fmt.Print("test1:")
    	print(myFlagSet.Lookup("test1"))
    	myFlagSet.Parse([]string{"-test=12345"})
    	fmt.Print("test:")
    	print(myFlagSet.Lookup("test"))
    	fmt.Print("test1:")
    	print(myFlagSet.Lookup("test1"))
    }
    
运行结果

    //  ./testflagsetlookup -test "12345"      
    test:default value
    test1:<nil>
    test:12345
    test1:<nil>
