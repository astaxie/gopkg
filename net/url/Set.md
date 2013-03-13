func (v Values) Set(key, value string)
#func (Values) Set
[参数列表]：

- key 键名

[返回值]：

- 

[功能说明]：

设置一个键关联的值，如该键存在，该键的值将被直接替换。如果该键不存在将创建

[代码实例]：
    package main
    
    import (
      "fmt"
    	"net/url"
    )
    func main() {
    	v := url.Values{}
    	v.Add("body", "good")
    	v.Add("body", "good1")
    	v.Add("body", "good2")
    	fmt.Println(v["body"])
    	//输出结果：[good good1 good2]
    	v.Set("body", "Sarah")
    	fmt.Println(v["body"])
    	//输出结果：[Sarah]
    }
