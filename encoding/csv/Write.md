## func (w *Writer) Write(record []string) (err error)

### 参数列表：

- record 将要输出的数据

### 返回值：

- err 输出时产生的错误

### 功能说明：

该函数将数据以 csv 的格式输出。

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
	writer.Flush() // 不要忘记调用 Flush 函数。
}

```