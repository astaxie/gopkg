func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

参数列表:

- v 序列化对象
- prefix 字首
- indent 缩减的字段

返回值:

- []byte 序列化的结果
- error 错误

功能说明:

这个函数主要是用于将对象v序列化为缩进格式json，[]byte,

代码实例:

    package main

	import (
		"encoding/json"
		"fmt"
		"os"
	)
	
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	
	func main() {
		group := ColorGroup{
			ID:     1,
			Name:   "Reds",
			Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		}
		b, err := json.MarshalIndent(group, "##", "**") //实际应用中更多的可能为 空格，这里为了便于观察
		if err != nil {
			fmt.Println("error:", err)
		}
		os.Stdout.Write(b)
	}




显示结果:

	{
	##**"ID": 1,
	##**"Name": "Reds",
	##**"Colors": [
	##****"Crimson",
	##****"Red",
	##****"Ruby",
	##****"Maroon"
	##**]
	##}