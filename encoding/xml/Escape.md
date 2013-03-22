##func Escape(w io.Writer, s []byte)

参数列表:

- w 输出接口，打印结果
- s 表示待转换的数据

返回值:

- 无

功能说明:

将s中的文本数据中的xml保留字符转换成预定义的实体引用，并输出到w。转换对照:

    "   &#34;
    '   &#39;
    &   &amp;
    <   &lt;
    >   &gt;
    
代码实例:

    package main
    
    import (
    	"encoding/xml"
    	"os"
    )
    
    func main() {
    	str := "\"'&<>"
    	xml.Escape(os.Stdout, []byte(str))
    }

输出结果:

    &#34;&#39;&amp;&lt;&gt;
