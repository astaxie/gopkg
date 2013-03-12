# type Replacer
这是一个字符串替换的对象

# func NewReplacer(oldnew ...string) *Replacer

参数列表

- oldnew是一个slice，是一个需要替换的字符串和新的字符串的配对出现

返回参数

- Replacer返回一个替换对象

Replacer方法列表

- func (r *Replacer) Replace(s string) string   // 把字符串替换为oldnew定义的
- func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)  //替换之后的字符串写入到w之中，返回写入的数量

应用示例，下面代码来自于beego的模板替换：

	package main

	import (
		"fmt"
		"strings"
		"os"
	)
	
	func main() {
		patterns := []string{"abc", "efg"}
		replacer := strings.NewReplacer(patterns...)
		format := replacer.Replace("abc is abc is abc")
		fmt.Println(format)
		//efg is efg is efg
		replacer.WriteString(os.Stdout,"abc is abc is abc")
		//efg is efg is efg
	}