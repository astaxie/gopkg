## func NewWriter(w io.Writer) *Writer

### 参数列表：

- r io.Writer 类型的输出对象

### 返回值：

- 返回一个 Writer 对象，将数据以 csv 格式输出。

### 功能说明：

该函数生成一个能够将数据以 csv 格式输出的 Writer 对象。

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