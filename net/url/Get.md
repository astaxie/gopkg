#func (Values) Get
[参数列表]：

- key 键名

[返回值]：

- 

[功能说明]：
Get获取一个键的值，如果该键有多个值将取第一个值。如果该键没有关联的值将返回空的字符串。若要访问多个值直接使用map

[代码实例]：
    package main
    
    import (
      "fmt"
    	"net/url"
    )
    func main() {
    	v := url.Values{}
    	v.Add("body", "good")
    	v.Add("friend", "Sarah")
    	v.Add("friend", "Zoe")
    	fmt.Println(v.Get("body"))
    	//输出结果：good
    	fmt.Println(v.Get("friend"))
    	//输出结果：Sarah  (由于friend键关联多个值，自动匹配第一个）
    	fmt.Println(v["friend"])
    	//输出结果：[Sarah Zoe] （friend键关联多个值，直接取map)
    }
