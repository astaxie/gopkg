# 包名

函数列表

- func After(d Duration) <-chan time
- func Sleep(d Duration)
- func Tick(d Duration) <-chan time

- func ParseDuration(s string) (Duration, err)
- func Since(t Time) Duration
- func (d Duration) Hours() float64
- func (d Duration) Minutes() float64
- func (d Duration) Nanoseconds() int64
- func (d Duration) Seconds() float64
- func (d Duration) String() string

- func FixedZone(name string, offset int) *Location
- func LoadLocation(name string) (*Location, error)
- func (l *Location) String() string

- func (m Month) String() string

- func (e *ParseError) Error() string

- func NewTicker(d Duration) *Ticker
- func (t *Ticker) Stop()

- func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
- func Now() Time
- func Parse(layout, value string) (Time, error)
- func Unix(sec int64, nsec int64) Time
- func (t Time) Add(d Duration) Time
- func (t TIme) AddDate(years int, months int, days int) Time
- func (t Time) After(u Time) bool
- func (t Time) Before(u Time) bool
- func (t Time) Clock() (hour, min, sec int)
- func (t TIme) Date() (year int, month Month, day int)
- func (t Time) Day() int
- func (t Time) Equal(u Time) bool
- func (t Time) Format(layout string) string
- func (t Time) GobDecode(buf []byte) error
- func (t Time) GobEncode() ([]byte, error)
- func (t Time) Hour() int
- func (t Time) ISOWeek() (year, week int)
- func (t Time) In(loc *Location) Time
- func (t Time) IsZero() bool
- func (t Time) Local() Time
- func (t Time) Location() *Location
- func (t Time) MarshalJSON() ([]byte, error)
- func (t Time) Minute() int
- func (t Time) Month() Month
- func (t Time) Nanosecond() int
- func (t Time) Second() int
- func (t Time) String() string
- func (t Time) Sub(u Time) Duration
- func (t Time) UTC() Time
- func (t Time) Unix() int64
- func (t Time) UnixNano() int64
- func (t Time) UnmarshalJSON(data []byte) (err error)
- func (t Time) Weekday() Weekday
- func (t Time) Year() int
- func (t Time) Zone() (name string, offset int)

- func AfterFunc(d Duration, f func()) *Timer
- func NewTimer(d Duration) *Timer
- func (t *Timer) Stop() (ok bool)

- func (d Weekday) String() string
