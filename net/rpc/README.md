# net/rpc

### CONSTANTS

    const (
        // Defaults used by HandleHTTP
        DefaultRPCPath   = "/_goRPC_"
        DefaultDebugPath = "/debug/rpc"
    )

### VARIABLES

    var DefaultServer = NewServer()
        DefaultServer is the default instance of *Server.
    var ErrShutdown = errors.New("connection is shut down")

- [func Accept(lis net.Listener)](Accept.md)
- [func HandleHTTP()](HandleHTTP.md)
- [func Register(rcvr interface{}) error](Register.md)
- [func RegisterName(name string, rcvr interface{}) error](RegisterName.md)
- [func ServeCodec(codec ServerCodec)](ServeCodec.md)
- [func ServeConn(conn io.ReadWriteCloser)](ServeConn.md)
- [func ServeRequest(codec ServerCodec) error](ServeRequest.md)

### type Call

    type Call struct {
        ServiceMethod string      // The name of the service and method to call.
        Args          interface{} // The argument to the function (*struct).
        Reply         interface{} // The reply from the function (*struct).
        Error         error       // After completion, the error status.
        Done          chan *Call  // Strobes when call is complete.
    }

### type Client

    type Client struct {
        // contains filtered or unexported fields
    }

- [func Dial(network, address string) (*Client, error)](Dial.md)
- [func DialHTTP(network, address string) (*Client, error)](DialHTTP.md)
- [func DialHTTPPath(network, address, path string) (*Client, error)](DialHTTPPath.md)
- [func NewClient(conn io.ReadWriteCloser) *Client](NewClient.md)
- [func NewClientWithCodec(codec ClientCodec) *Client](NewClientWithCodec.md)
- [func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error](Client_Call.md)
- [func (client *Client) Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call](Client_Go.md)

### type ClientCodec

    type ClientCodec interface {
        WriteRequest(*Request, interface{}) error
        ReadResponseHeader(*Response) error
        ReadResponseBody(interface{}) error

        Close() error
    }

### type Request

    type Request struct {
        ServiceMethod string // format: "Service.Method"
        Seq           uint64 // sequence number chosen by client
        // contains filtered or unexported fields
    }

### type Response

    type Response struct {
        ServiceMethod string // echoes that of the Request
        Seq           uint64 // echoes that of the request
        Error         string // error, if any.
        // contains filtered or unexported fields
    }

### type Server
    type Server struct {
        // contains filtered or unexported fields
    }

- [func NewServer() *Server](NewServer.md)
- [func (server *Server) Accept(lis net.Listener)](Server_Accept.md)
- [func (server *Server) HandleHTTP(rpcPath, debugPath string)](Server_HandleHTTP.md)
- [func (server *Server) Register(rcvr interface{}) error](Server_Register.md)
- [func (server *Server) RegisterName(name string, rcvr interface{}) errorr](Server_RegisterName.md)
- [func (server *Server) ServeCodec(codec ServerCodec)](Server_ServeCodec.md)
- [func (server *Server) ServeConn(conn io.ReadWriteCloser)](Server_ServeConn.md)
- [func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request)]
- [func (server *Server) ServeRequest(codec ServerCodec) error](Server_ServeRequest.md)

### type ServerCodec
    type ServerCodec interface {
        ReadRequestHeader(*Request) error
        ReadRequestBody(interface{}) error
        WriteResponse(*Response, interface{}) error
        Close() error
    }

### type ServerError
    type ServerError string

- [func (e ServerError) Error() string](ServerError_Error.md)
