## func (r *Reader) ReadAll() (records [][]string, err error)

### 参数列表：

### 返回值：

- records []][]string 类型，对应于 csv 文件的多行数据。
- err 读取时产生的错误

### 功能说明：

该函数用于一次性读取所有未读的多行数据。

### 代码实例：

```go
package main

import (
	"encoding/csv"
	"fmt"
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
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(records)
}

```