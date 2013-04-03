## func Join(elem ...string) string

参数列表

- elem 表示要连接的路径


返回值：

- 返回string

功能说明：

这个函数主要是连接路径，返回的结果是已经Clean的，如果是空路径就忽略



代码实例：
~~~
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Join("a", "b", "c")) // a/b/c
	fmt.Println(path.Join("a", "", "c"))  // a/c
	fmt.Println(path.Join("a", "../bb/../c", "c")) // c/c
}
~~~