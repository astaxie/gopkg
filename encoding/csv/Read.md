## func (r *Reader) Read() (record []string, err error)

### 参数列表：

### 返回值：

- record 字符串类型的 slice ，对应于 csv 文件的一行数据。
- err 读取时产生的错误

### 功能说明：

该函数用于读取一行数据。当读取完所有数据后，继续调用 `Read()` ，产生的是 EOF 错误，需要与因原始数据的格式问题导致产生 ParseError 的错误区分开来。

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