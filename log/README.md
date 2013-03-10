# log包

函数列表

Constants
- [func Fatal(v ...interface{})](Fatal.md)
- [func Fatalf(format string, v ...interface{})](Fatalf.md)
- [func Fatalln(v ...interface{})](Fatalln.md)
- [func Flags() int](Flags.md)
- [func Panic(v ...interface{})](Panic.md)
- [func Panicf(format string, v ...interface{})](Panicf.md)
- [func Panicln(v ...interface{})](Panicln.md)
- [func Prefix() string](Prefix.md)
- [func Print(v ...interface{})](Print.md)
- [func Printf(format string, v ...interface{})](Printf.md)
- [func Println(v ...interface{})](Println.md)
- [func SetFlags(flag int)](SetFlags.md)
- [func SetOutput(w io.Writer)](SetOutput.md)
- [func SetPrefix(prefix string)](SetPrefix.md)
type Logger
	- [func New(out io.Writer, prefix string, flag int) *Logger](New.md)
    - func (l *Logger) Fatal(v ...interface{})				参考：[func Fatal(v ...interface{})](Fatal.md)
    - func (l *Logger) Fatalf(format string, v ...interface{})				参考：[func Fatalf(format string, v ...interface{})](Fatalf.md)
    - func (l *Logger) Fatalln(v ...interface{})				参考：[func Fatalln(v ...interface{})](Fatalln.md)
    - func (l *Logger) Flags() int				参考：[func Flags() int](Flags.md)
    - [func (l *Logger) Output(calldepth int, s string) error](Output.md)
    - func (l *Logger) Panic(v ...interface{})				参考：[func Panic(v ...interface{})](Panic.md)
    - func (l *Logger) Panicf(format string, v ...interface{})				参考：[func Panicf(format string, v ...interface{})](Panicf.md)
    - func (l *Logger) Panicln(v ...interface{})				参考：[func Panicln(v ...interface{})](Panicln.md)
    - func (l *Logger) Prefix() string				参考：[func Prefix() string](Prefix.md)
    - func (l *Logger) Print(v ...interface{})				参考：[func Print(v ...interface{})](Print.md)
    - func (l *Logger) Printf(format string, v ...interface{})				参考：[func Printf(format string, v ...interface{})](Printf.md)
    - func (l *Logger) Println(v ...interface{})				参考：[func Println(v ...interface{})](Println.md)
    - func (l *Logger) SetFlags(flag int)				参考：[func SetFlags(flag int)](SetFlags.md)
    - func (l *Logger) SetPrefix(prefix string)				参考：[func SetPrefix(prefix string)](SetPrefix.md)
 
