## func Match(pattern, name string) (matched bool, err error)

参数列表

- pattern 匹配的模式
- name 原始字符串

返回值：

- 返回bool error

功能说明：

这个函数主要是文件名匹配，只有完全匹配则返回true，nil
~~~
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Match("*","alll")) //true nil
	fmt.Println(path.Match("*","a/lll")) //false nil
	fmt.Println(path.Match("?","alll")) //false nil
	fmt.Println(path.Match("?","a")) //true nil
	fmt.Println(path.Match("a","a")) //true nil
	fmt.Println(path.Match("\\a","a")) //true nil
}
~~~