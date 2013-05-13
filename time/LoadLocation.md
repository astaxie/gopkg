# func LoadLocation(name string) (*Location error)

参数列表：

- name 时区名称

返回值：

- Location
- error

功能说明：

LoadLocation返回name对应的Location。

如果name是""或“UTC”,LoadLocation返回UTC，如果name是“Local”， LoadLocation返回Local。

否则，name取为IANA时区数据库中的一个文件对应的位置名称，如“America/New_York”。

LoadLocation使用的时区数据库可能不会被所有的系统所提供，尤其是非Unix系统。LoadLocation在ZONEINFO环境变量指定的目录或者解压的zip文件中寻找，如果没有，在Unix系统中已知的安装位置寻找，最后在 $GOROOT/lib/time/zoneinfo.zip中寻找

代码实例：

	package main
	
	import (
	    "time"
	    "fmt"
	)
	
	func main() {
	    loc, err := time.LoadLocation("America/New_York")
	    if err != nil {
	        fmt.Println(err)
	        return
	    }
	
	    fmt.Println(loc)
	}
