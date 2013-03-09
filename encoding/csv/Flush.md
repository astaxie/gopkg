## func (w *Writer) Flush()

### 参数列表：

### 返回值：

### 功能说明：

该函数将缓冲内的数据写入真实文件。

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

	writer := csv.NewWriter(file)
	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	writer.Flush()
}

```