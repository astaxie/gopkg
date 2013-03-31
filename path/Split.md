## func Split(path string) (dir, file string)

参数列表

- path 表示路径字符串


返回值：

- 返回string,string

功能说明：

这个函数主要是分离路径中的文件目录和文件



代码实例：
~~~
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Split("static/myfile.css")) // static/ myfile.css
	fmt.Println(path.Split("static"))	// "" static
}
~~~