## func After(d Duration) <-chan Time
参数列表
- d 时间跨度
返回值
- 返回Time类型chan
功能说明：
这个函数用来实现指定时间的延时，返回一个单向的读取chan
代码实例：
```go
package main

import (
  "fmt"
	"time"
)

const dur = 3

func main() {
	fmt.Println("Hello, playground")
	<-time.After(dur * time.Second)
	fmt.Println(dur, "seconds later")
}
```
