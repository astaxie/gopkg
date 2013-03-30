# 常量

常量列表：

~~~
- Ldate         日期 年/月/日
- Ltime         时间 时:分:秒
- Lmicroseconds 时间 .毫秒于Ltime之后
- Llongfile     完整文件名:行号
- Lshortfile    文件名，此标志位优先于 Llongfile
- LstdFlags     = Ldate 并且 Ltime
~~~

功能说明：

标志位常量控制日志格式。

代码实例：
```go
package main

import "log"

func main() {
  log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("log:")
}
```

输出：
~~~
2009/11/10 23:00:00 log.go:7: log:
~~~
