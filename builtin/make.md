## func make(Type, size IntegerType) Type

参数列表：
- Type 类型 slice,map,chan 之一
- size IntegerType 
返回值：
- Type
功能说明：

make 内建函数分配并初始化一个类型为切片、映射、或（仅仅为）信道的对象。 与 new 相同的是，其第一个实参为类型，而非值。不同的是，make 的返回类型 与其参数相同，而非指向它的指针。其具体结果取决于具体的类型：

~~~
slice：
	size 指定了切片长度。该切片的容量等于其长度。
	第二个整数实参可用来指定不同的容量；
	它必须不小于其长度，因此 make([]int, 0, 10) 
	会分配一个长度为0，	容量为10的切片。
map：
	初始分配取决于 size，但产生的 map 长度为0。
	size 可以省略，这种情况下就会分配一个最小的起始容量。
chan：
	信道的缓存根据指定的缓存容量初始化。
	若 size 为零或被省略，该信道即为无缓存的。
~~~

代码实例：

~~~go
package main

import "fmt"

func main() {
    a := make([]int, 5)
    printSlice("a", a)
    b := make([]int, 0, 5)
    printSlice("b", b)
    c := b[:2]
    printSlice("c", c)
    d := c[2:5]
    printSlice("d", d)
}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n",
        s, len(x), cap(x), x)
}
~~~

输出：

~~~
a len=5 cap=5 [0 0 0 0 0]
b len=0 cap=5 []
c len=2 cap=5 [0 0]
d len=3 cap=3 [0 0 0]
~~~