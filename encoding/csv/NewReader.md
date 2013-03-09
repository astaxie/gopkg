## func NewReader(r io.Reader) *Reader

### 参数列表：

- r io.Reader 类型的输入对象

### 返回值：

- 返回能够解析 csv 格式输入的 Reader 对象

### 功能说明：

该函数返回一个能够解析 csv 格式的 Reader 对象。

### 代码实例：

```go
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("demo.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF { // EOF 表示已经读取了所有的数据
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(record) // record has the type []string

	}
}

```