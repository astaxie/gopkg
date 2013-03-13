#func (Values) Add
[参数列表]：

- key 键名称
- value 值类容(字符串）

[返回值]：

- v 返回map

[功能说明]：
给某个键(key)增加值，它将追加到键(key)现有值的后面;注意键(key)区分大小写，如键不存在将创建

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
    	fmt.Println(v)
    	//输出结果：map[friend:[aaaa bbbb cccc] good:[baby]]
    }
