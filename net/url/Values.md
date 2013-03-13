#type Values

[参数列表]：

- 

[返回值]：

-

[功能说明]：
maps值与键的操作，它通常用于查询参数或值，不像在http.Header的map,在maps中的键是区分大小写的

[代码实例]：

    package main    
    import (
      "fmt"
    	"net/url"
    )
    
    func main() {
    	v := url.Values{}
    	v.Set("name", "Ava")
    	v.Add("friend", "Jess")
    	v.Add("friend", "Sarah")
    	v.Add("friend", "Zoe")
    	// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
    	fmt.Println(v.Get("name"))
    	fmt.Println(v.Get("friend"))
    	fmt.Println(v["friend"])
    	fmt.Println(v)
    }
    //输出结果：
    //Ava
    //Jess
    //[Jess Sarah Zoe]
    //map[friend:[Jess Sarah Zoe] name:[Ava]]
