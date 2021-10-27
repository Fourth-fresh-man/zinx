package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

// Server iServer的接口实现，定义一个Server的服务器模块
type Server struct {
	// 服务器名
	name string
	// 服务器IP版本
	IPVersion string
	// 服务器IP
	IP string
	// 服务器port
	Port int
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP:%s, Port:%d", s.IP, s.Port)
	fmt.Println("\nStarting...")

	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err!=nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}
		Listener, err := net.ListenTCP(s.IPVersion, addr)
		if err!=nil {
			fmt.Println("listen",s.IPVersion,"err",err)
			return
		}
		fmt.Println("start Zinx server succ,",s.name)
		fmt.Println("Listening.....")
		for {
			conn, err := Listener.AcceptTCP()
			if err!=nil {
				fmt.Println("Accept err", err)
				continue
			}

			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err!=nil {
						fmt.Println("recv buf err", err)
						continue
					}

					fmt.Printf("recv client content=%s cnt=%d\n", buf, cnt)
					if _, err := conn.Write(buf[:cnt]); err!=nil {
						fmt.Println("write back err", err)
						continue
					}
				}
			}()
		}
	}()

}
func (s *Server) Stop() {
	// todo 停止服务器，并回收资源和一些额外业务

}
func (s *Server) Serve() {
	s.Start()

	// todo 启动服务器后的额外业务

	// 阻塞
	select {}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		name: name,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 8999,
	}
	return s
}