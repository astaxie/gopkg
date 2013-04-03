func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error

参数列表:

- dst 表示字符缓冲区指针
- src 表示JSON格式的字符串切片
- prefix 字首
- indent 缩减的字段

返回值:

- error

功能说明:

这个函数主要是用于将src以缩进形式添加到缩进的JSON内容dst中, prefix用于填充字首，indent用于填充缩进。

代码实例:

    package main

	import (
		"bytes"
		"encoding/json"
		"fmt"
	)
	
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	
	func main() {
		dst := new(bytes.Buffer)
		src := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
		json.Indent(dst, src, "##", "**") //实际应用中，可能更多的是使用空格来作为prefix和indent，这里主要为了方便观察
		json.Indent(dst, src, "##", "**") //实际应用中，可能更多的是使用空格来作为prefix和indent，这里主要为了方便观察
		fmt.Println(dst)
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
	##}{
	##**"ID": 1,
	##**"Name": "Reds",
	##**"Colors": [
	##****"Crimson",
	##****"Red",
	##****"Ruby",
	##****"Maroon"
	##**]
	##}