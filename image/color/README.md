# imgage/color

函数列表:
- RGBToYCbCr
- YCbCrToRGB
   
## func RGBToYCbCr(r, g, b uint8) (uint8, uint8, uint8) 

参数列表:
- r rgb颜色的红色分量
- g rgb颜色的green 分量
- b rgb颜色的blue 分量

返回值:
- Y
- Cb
- Cr 

功能说明:
	RGB与YCbCr颜色空间转换
	YCbCr参考链接:
		http://zh.wikipedia.org/zh-cn/YCbCr
		http://www.cnblogs.com/starspace/archive/2008/12/16/1356007.html
		
	// RGB转换为YCbCr
	// The JFIF specification says:
	//	Y' =  0.2990*R + 0.5870*G + 0.1140*B
	//	Cb = -0.1687*R - 0.3313*G + 0.5000*B + 128
	//	Cr =  0.5000*R - 0.4187*G - 0.0813*B + 128
	// http://www.w3.org/Graphics/JPEG/jfif3.pdf says Y but means Y'.

	
// YCbCr与RGB颜色空间转换
func YCbCrToRGB(y, cb, cr uint8) (uint8, uint8, uint8) 

	
代码实例：

```go
package main

import (
	"image/color"
)

func main() {
	var r, g, b uint8

	r, g, b = 255, 255, 0
	y, cb, cr := color.RGBToYCbCr(r, g, b)
	println("RGBToYCbCr:y,ca,cb:", y, cb, cr)
	//y,ca,cb: 226 1 149
	r, g, b = color.YCbCrToRGB(y, cb, cr)
	println("YCbCrToRGB:", r, g, b)
	//YCbCrToRGB: 255 255 1
}

```


接口 type Color interface {
	//返回 red green blue 透明分量
	RGBA() (r, g, b, a uint32)
}

结构体RGBA 
type RGBA struct {
	//存储 red green blue 透明分量
	R, G, B, A uint8
}

函数
func (c RGBA) RGBA() (r, g, b, a uint32) 

功能说明:
	对当前颜色 返回Rgba 分量,每个分量是32位无符号整数
		每个分量是当前分量左移8位在并上当前分量
		比如:r = uint32(c.R)
			r |= r << 8
			
		 r, g, b, a uint32 
		 
示例代码:
```go
package main

import (
	"fmt"
	"image/color"
)

func main() {
	c := color.RGBA{255, 6, 1, 0}
	fmt.Println(c)
	//{255 6 1 0}

	r, g, b, a := c.RGBA()
	fmt.Println("c.RGBA", r, g, b, a)
	//c.RGBA 65535 1542 257 0

	fmt.Printf("c.RGBA bin %b,%b,%b,%b", r, g, b, a)
	//c.RGBA bin 1111111111111111,11000000110,100000001,0%
}
```

