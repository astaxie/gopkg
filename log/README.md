# log包

常量

- [Constants](Constants.md)

函数列表

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

- type Logger
 - [func New(out io.Writer, prefix string, flag int) *Logger](New.md)
 - [func (l *Logger) Fatal(v ...interface{})](lFatal.md)
 - [func (l *Logger) Fatalf(format string, v ...interface{})](lFatalf.md)
 - [func (l *Logger) Fatalln(v ...interface{})](lFatalln.md)
 - [func (l *Logger) Flags() int](lFlags.md)
 - [func (l *Logger) Output(calldepth int, s string) error](Output.md)
 - [func (l *Logger) Panic(v ...interface{})](lPanic.md)
 - [func (l *Logger) Panicf(format string, v ...interface{})](lPanicf.md)
 - [func (l *Logger) Panicln(v ...interface{})](lPanicln.md)
 - [func (l *Logger) Prefix() string](lPrefix.md)
 - [func (l *Logger) Print(v ...interface{})](lPrint.md)
 - [func (l *Logger) Printf(format string, v ...interface{})](lPrintf.md)
 - [func (l *Logger) Println(v ...interface{})](lPrintln.md)
 - [func (l *Logger) SetFlags(flag int)](lSetFlags.md)
 - [func (l *Logger) SetPrefix(prefix string)](lSetPrefix.md)
 
