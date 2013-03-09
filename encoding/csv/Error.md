## func (e *ParseError) Error() string

### 参数列表：

### 返回值：

- 返回字符串

### 功能说明：

这个函数主要用来返回 ParseError 实例的具体错误信息。

### 代码实例：

```go
package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	reader := csv.NewReader(strings.NewReader("cell10,cell11\ncell20"))
	_, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
	}
}

```

运行后，终端将输出 `line 2, column 0: wrong number of fields in line` .