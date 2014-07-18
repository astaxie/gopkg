# func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)

参数列表

- s         转义后的字符串
- quote     字符串使用的引号符。
    如果设置为单引号，则 s 中允许出现 \' 字符，不允许出现单独的 ' 字符
    如果设置为双引号，则 s 中允许出现 \" 字符，不允许出现单独的 " 字符
    如果设置为 0，则不允许出现 \' 或 \" 字符，可以出现单独的 ' 或 " 字符

返回值：

- value     解码后的字符
- multibyte value是否为多字节字符
- tail      字符串 s 除去 value 后的剩余部分
- err       返回 s 中是否存在语法错误

功能说明：

- 将 s 中的第一个字符“取消转义”并解码


代码实例：

    func main() {
        sr := `\"大\\家\\好！\"`
        var c rune
        var mb bool
        var err error
        for ; len(sr) > 0; c, mb, sr, err = strconv.UnquoteChar(sr, '"') {
            fmt.Println(c, mb,sr,err)
        }
    }

代码输出：

    0 false \"大\\家\\好！\" <nil>
    34 false 大\\家\\好！\" <nil>
    22823 true \\家\\好！\" <nil>
    92 false 家\\好！\" <nil>
    23478 true \\好！\" <nil>
    92 false 好！\" <nil>
    22909 true ！\" <nil>
    65281 true \" <nil>