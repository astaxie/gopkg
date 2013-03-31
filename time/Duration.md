# type Duration int64

以int64表示两个时间点之间的间隔（纳秒数），最多表示290年

常量

- Nanosecond Duration 	= 1
- Microsecond			= 1000 * Nanosecond
- Millissecond			= 1000 * Microsecond
- Second				= 1000 * Millisecond
- Minute				= 60 * Second
- Hour					= 60 * Minute

通常没有天或者更大的时间单位，是为了避免混淆跨越夏令时区转换。
