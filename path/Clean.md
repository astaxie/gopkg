## func Clean(path string) string

参数列表

- path 表示需要清理的路径字符串 


返回值：

- 返回string

功能说明：

这个函数主要是返回等价的最短路径
清理规则：
-		1.用一个斜线替换多个斜线
-		2.清除当前路径.
-		3.清除内部的..和他前面的元素，如a/b/.. 得到结果a
-		4.以/..开头的，变成/


代码实例：
~~~
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Clean("a/c"))// a/c
	fmt.Println(path.Clean("a//c"))//以一个/代替// , a/c
	fmt.Println(path.Clean("a/c/."))//清除. , a/c
	fmt.Println(path.Clean("a/c/b/.."))// 清除内部..以及前面的元素b; a/c
	fmt.Println(path.Clean("111/../a/c"))// 清除内部..以及前面的元素111; a/c
	fmt.Println(path.Clean("/../a/b/../././/c"))// 清除/..开头,..以及前面的元素b; /a/c
}
~~~