## func IsAbs(path string) bool

参数列表

- path 表示路径字符串 


返回值：

- 返回bool

功能说明：

这个函数主要是判断路径是不是绝对路径，如果是绝对路径返回true



代码实例：
~~~
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.IsAbs("/home/zzz/go.pdf")) // true
	fmt.Println(path.IsAbs("home/zzz/go.pdf"))  // false
}
~~~