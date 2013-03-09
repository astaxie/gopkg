## func DetectContentType(data []byte) string

参数列表

- data 需要检测类型的数据 

返回值：

- 返回检测好的MIME类型。如果不能检测到任何MIME类型，返回"application/octet-stream"。

功能说明：

该函数实现了一个算法，用来检测指定的数据是否符合http://mimesniff.spec.whatwg.org/所定义的MIME类型。
该函数最多需要数据的前512个字节。
该函数总是会返回一个有效的MIME类型。
如果它不能够识别数据，将会返回"application/octet-stream"。

代码实例

  package main
	import (
		"fmt"
		"net/http"
	)

	func main() {
		var notingData []byte = []byte{0x01, 0x02}
		var gifData []byte = []byte{0x01, 0x02}				
		fmt.Println(http.DetectContentType(notingData))
		fmt.Println(http.DetectContentType(gifData))
	}


