# type Month int

一年中的一个月

常量：

- January Month = 1 + iota
- February		
- March
- April
- May
- June
- July
- August
- September
- October
- November
- December

代码实例：

	package main
	
	import (
		"fmt"
		"time"
	)
	
	func main() {
		_, m, _ := time.Now().Date()
		fmt.Println(m)
	}