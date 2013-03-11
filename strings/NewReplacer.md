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

	// Date takes a PHP like date func to Go's time fomate
	func Date(t time.Time, format string) (datestring string) {
		patterns := []string{
			// year
			"Y", "2006", // A full numeric representation of a year, 4 digits	Examples: 1999 or 2003
			"y", "06", //A two digit representation of a year	Examples: 99 or 03
	
			// month
			"m", "01", // Numeric representation of a month, with leading zeros	01 through 12
			"n", "1", // Numeric representation of a month, without leading zeros	1 through 12
			"M", "Jan", // A short textual representation of a month, three letters	Jan through Dec
			"F", "January", // A full textual representation of a month, such as January or March	January through December
	
			// day
			"d", "02", // Day of the month, 2 digits with leading zeros	01 to 31
			"j", "2", // Day of the month without leading zeros	1 to 31
	
			// week
			"D", "Mon", // A textual representation of a day, three letters	Mon through Sun
			"l", "Monday", // A full textual representation of the day of the week	Sunday through Saturday
	
			// time
			"g", "3", // 12-hour format of an hour without leading zeros	1 through 12
			"G", "15", // 24-hour format of an hour without leading zeros	0 through 23
			"h", "03", // 12-hour format of an hour with leading zeros	01 through 12
			"H", "15", // 24-hour format of an hour with leading zeros	00 through 23
	
			"a", "pm", // Lowercase Ante meridiem and Post meridiem	am or pm
			"A", "PM", // Uppercase Ante meridiem and Post meridiem	AM or PM
	
			"i", "04", // Minutes with leading zeros	00 to 59
			"s", "05", // Seconds, with leading zeros	00 through 59
		}
		replacer := strings.NewReplacer(patterns...)
		format = replacer.Replace(format)
		datestring = t.Format(format)
		return
	}