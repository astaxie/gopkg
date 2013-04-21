## func SplitList(path string) []string  

参数列表:

- 路径字符串
 
返回值:

- 路径列表。

功能说明:

将路径字符串使用路径列表分隔符分开。路径列表分隔符见os.PathListSeparator, linux的
路径列表分隔符是":", windows的路径列表分隔符是";"，主要用于PATH或是GOPATH等环境变
量。

示例：

~~~
package main

import (
    "fmt"
    "path/filepath"
)

func main(){
    fmt.Printf("%v\n", filepath.SplitList("a/b;c/d"))  //[a/b; c/d]
}
~~~

