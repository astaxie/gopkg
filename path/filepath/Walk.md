## func Walk(root string, walkFunc WalkFunc) error

参数列表:

- root: 需要遍历的目录
- walkFunc: 处理单个文件的函数(linux下目录和文件统称为文件),类型为WalkFunc, WalkFunc的定义为type WalkFunc func(path string, info os.FileInfo, err error) error。path为文件名，info文件的信息，遍历中的错误。 

返回值列表:

- 返回error


功能说明:
使用walkFunc遍历处理path目录下的所的文件，包括path。

示例:

~~~
package main

import (
    "os"
    "fmt"
    "path/filepath"
)

func walkFunc (path string, info os.FileInfo, err error) error {
    fmt.Printf("%s\n", path)
    return nil
}

func main(){
    //遍历打印所有的文件名
    filepath.Walk("/home", walkFunc)
}
~~~

