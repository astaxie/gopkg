# os/exec函数列表

- [func LookPath(file string) (string, error)](LookPath.md)
- [func Command(name string, arg ...string) *Cmd](Command.md)
- [func (c *Cmd) CombinedOutput() ([]byte, error)](CombinedOutput.md)
- [func (c *Cmd) Output() ([]byte, error)](Output.md)
- [func (c *Cmd) Run() error](Run.md)
- [func (c *Cmd) Start() error](Start.md)
- [func (c *Cmd) StderrPipe() (io.ReadCloser, error)](StderrPipe.md)
- [func (c *Cmd) StdinPipe() (io.WriteCloser, error)](StdinPipe.md)
- [func (c *Cmd) StdoutPipe() (io.ReadCloser, error)](StdoutPipe.md)
- [func (c *Cmd) Wait() error](Wait.md)
- [func (e *Error) Error() string](Error.md)
- [func (e *ExitError) Error() string](ExitError.md)
