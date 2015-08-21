package tcpproxy

type TCPServer struct{
	Debug bool
	Server
	process *Process
}

// Connection 连接
type Connection struct{

}
// Connections 连接表
type Connections []Connection

// Pool 资源池
type Pool struct{

}



