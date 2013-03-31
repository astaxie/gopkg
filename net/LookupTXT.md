## func LookupTXT(name string) (txt []string, err error)

参数列表:

- name 域名

返回列表:

- txt txt记录名
- err 错误信息

查找该域名下的txt记录

[百度百科-TXT记录](http://baike.baidu.com/view/1372987.htm)

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		txts,err:=net.LookupTXT("163.com")
		if err != nil {
			fmt.Println(err)
		}
		for _,txt := range txts {
			fmt.Println(txt) // 返回类似 v=spf1 include:spf.163.com -all
		}
	}