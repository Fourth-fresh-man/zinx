package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
	Conn *net.TCPConn
	ConnID uint32
	isClosed bool
	handleAPI ziface.HandleFunc
	ExitChan chan bool
}

func (c Connection) Start() {
	fmt.Println("Conn start...ConnID=", c.ConnID)
	defer fmt.Println("connID =", c.ConnID, "Reader is exit, remote addr is", c.RemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt,err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}
		if err = c.handleAPI(c.Conn, buf, cnt); err!=nil {
			fmt.Println(c.ConnID, "handle is error ", err)
			break
		}
	}
}

func (c Connection) Stop() {
	fmt.Println("Conn stop...ConnID=", c.ConnID)
	if c.isClosed {
		return
	}
	c.isClosed=true
	c.Conn.Close()
	close(c.ExitChan)
}

func (c Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c Connection) Send(data []byte) error {
	return nil
}

// NewConnection 初始连接模块
func NewConnection(conn *net.TCPConn, connID uint32, callbackApi ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn: conn,
		ConnID: connID,
		handleAPI: callbackApi,
		isClosed: false,
		ExitChan: make(chan bool,1),
	}
	return c
}
