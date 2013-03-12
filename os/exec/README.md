# os/exec函数列表

- func LookPath(file string) (string, error)
- func Command(name string, arg ...string) *Cmd
- func (c *Cmd) CombinedOutput() ([]byte, error)
- func (c *Cmd) Output() ([]byte, error)
- func (c *Cmd) Run() error
- func (c *Cmd) Start() error
- func (c *Cmd) StderrPipe() (io.ReadCloser, error)
- func (c *Cmd) StdinPipe() (io.WriteCloser, error)
- func (c *Cmd) StdoutPipe() (io.ReadCloser, error)
- func (c *Cmd) Wait() error
- func (e *Error) Error() string
- func (e *ExitError) Error() string
