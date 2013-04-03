## func (w *Writer) WriteAll(records [][]string) (err error)

### 参数列表：

- records 将要输出的数据

### 返回值：

- err 输出时产生的错误

### 功能说明：

该函数将(多行)数据以 csv 的格式输出。`WriteAll` 通过调用 `Write` 函数输出每行数据，最后调用 `Flush` 函数。

### 代码实例：

```go
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("demo.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	records := [][]string{{"10", "11", "12"}, {"20", "21", "22"}, {"30", "31", "32"}}

	err = writer.WriteAll(records)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}

```