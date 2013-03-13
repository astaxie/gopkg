#func (Values) Del
[参数列表]：

- key 键名称

[返回值]：

- v 返回map

[功能说明]：
删除某个键，同时将删除它所关联的值

[代码实例]：
    package main
    
    import (
      "fmt"
    	"net/url"
    )
    
    func main() {
    	v := url.Values{}
    	v.Add("good", "baby")
    	v.Add("friend", "aaaa")
    	v.Add("friend", "bbbb")
    	v.Add("friend", "cccc")
    	/////////
    	v.Del("friend")
    	fmt.Println(v)
    	//输出结果：map[good:[baby]]
    }
