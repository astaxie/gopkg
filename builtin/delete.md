## func delete(m map[Type]Type1, key Type)

参数列表：

- m map
- key 键

返回值：

- 无

功能说明：

delete 内建函数按照指定的键将元素从映射中删除。如果元素不存在，delete 为空操作。如果 m 为 nil，delete 引发 panic 错误。

代码实例：

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var nilmap map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	delete(m, "key")
	fmt.Println(m)
	delete(m, "Bell Labs")
	fmt.Println(m)

	delete(nilmap, "key")

}
```

输出:

~~~
map[Bell Labs:{40.68433 -74.39967}]
map[]
panic: runtime error: deletion of entry in nil map

goroutine 1 [running]:
main.main()
	/tmpfs/gosandbox-ffb3cedf_c062e804_9b5b7ada_f200355f_78578d79/prog.go:23 +0x1e0
~~~