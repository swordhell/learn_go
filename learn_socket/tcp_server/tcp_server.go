package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
)

type server struct {
	sync.Mutex
	ln       net.Listener
	SigClose chan bool
}

type session struct {
	conn     net.Conn
	SigClose chan bool
}

func (this *server) Listen() {
	ln, err := net.Listen("tcp4", "0.0.0.0:8089")
	if err != nil {
		fmt.Println("listen fail: ", err)
		return
	}
	this.ln = ln
	go func(l_this *server) {
		for {
			conn, cn_err := l_this.ln.Accept()
			if cn_err != nil {
				fmt.Println("accept fail", cn_err)
			}
			c := &session{conn: conn, SigClose: l_this.SigClose}
			c.StartSession()
		}
	}(this)
}

func (this *server) Close() {
	this.ln.Close()
}

func main() {
	sigClose := make(chan bool)
	svr := &server{SigClose: sigClose}
	svr.Listen()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c

	sigClose <- true
	svr.Close()
}

func (this *session) StartSession() {
	go func(l_this *session) {
		l_this.recvData()
	}(this)
}

func (this *session) recvData() {
	conn := this.conn
	for {
		data, err := procData(conn)
		if err != nil {
			fmt.Errorf("procData fail, err: %v\n", err)
			break
		}
		fmt.Printf("handleData addr: %v, data_size: %v",
			conn.RemoteAddr().String(), len(data))
		handleData(conn, data)
	}

}

func procData(conn net.Conn) ([]byte, error) {
	// 读取数据包头
	var bufHeader []byte = make([]byte, 2, 2)
	if _, err := io.ReadFull(conn, bufHeader); err != nil {
		return nil, err
	}

	var msgLen uint32
	msgLen = uint32(binary.LittleEndian.Uint16(bufHeader))

	msgData := make([]byte, msgLen)
	if _, err := io.ReadFull(conn, msgData); err != nil {
		return nil, err
	}

	return msgData, nil
}

func sendData(con net.Conn) {

}

func handleData(conn net.Conn, data []byte) {

}
