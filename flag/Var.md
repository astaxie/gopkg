## func Var(value Value, name string, usage string)

参数列表
- value Value   需要绑定的变量的名称
- name string 变量默认值
- usage string 提示信息

返回值

功能说明
- 将flag绑定到自定义变量上，自定义变量需要实现Value接口

代码示例
    
    package main
    
    import (
    	"errors"
    	"flag"
    	"fmt"
    	"strings"
    	"time"
    )
    
    type interval []time.Duration
    
    //实现String接口
    func (i *interval) String() string {
    	return fmt.Sprintf("%v", *i)
    }
    
    //实现Set接口,Set接口决定了如何解析flag的值
    func (i *interval) Set(value string) error {
    	//此处决定命令行是否可以设置多次-deltaT
    	if len(*i) > 0 {
    		return errors.New("interval flag already set")
    	}
    	for _, dt := range strings.Split(value, ",") {
    		duration, err := time.ParseDuration(dt)
    		if err != nil {
    			return err
    		}
    		*i = append(*i, duration)
    	}
    	return nil
    }
    
    var intervalFlag interval
    
    func init() {
    	flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
    }
    
    func main() {
    	flag.Parse()
    	fmt.Println(intervalFlag)
    }

运行结果
    
    //./testvar -deltaT 61m,72h,80s
    [1h1m0s 72h0m0s 1m20s]