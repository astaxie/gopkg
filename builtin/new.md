## func new(Type) *Type

参数列表：

- Type

返回值：

- *Type

功能说明：

new 内建函数分配内存。 其第一个实参为类型，而非值，其返回值为指向该类型新分配的零值的指针。

代码实例：

~~~go
package main

import "fmt"

type Vertex struct {
    X, Y int
}

func main() {
    var d Vertex
    v := new(Vertex)
    d.X =8
    fmt.Println(v,d)
    v.X, v.Y = 11, 9
    fmt.Println(v)
}
~~~

输出：

~~~
&{0 0} {8 0}
&{11 9}
~~~