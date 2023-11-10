package main
import "fmt"
import "net"
type Server struct {
	Ip string
	Port int
}

//创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip: ip,
		Port: port,
	}

	return server
}

func (this *Server) Handle(conn net.Conn){
	fmt.Println("连接已建立")
}
// 启动服务器接口
func (this *Server) Start() {
	// socket listen
  listener,_ :=	net.Listen("tcp",fmt.Sprintf("%s:%d",this.Ip,this.Port))
	defer listener.close();
	for{
		conn,_ := listener.Accept();
		go this.Handle(conn)
	}
	
}
