package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"
)

type session struct {
	conn      net.Conn
	writeChan chan []byte
}

type SessionSet map[*session]struct{}

type server struct {
	tempDelay time.Duration
	sync.Mutex
	ln         net.Listener
	wgLn       sync.WaitGroup
	wgSession  sync.WaitGroup
	sessionSet SessionSet
}

func (this *server) Close() {
	this.ln.Close()
	this.wgLn.Wait()
}

func (this *server) ProcTempErr() {
	tempDelay := this.tempDelay
	if tempDelay == 0 {
		tempDelay = 5 * time.Millisecond
	} else {
		tempDelay *= 2
	}
	if max := 1 * time.Second; tempDelay > max {
		tempDelay = max
	}
	fmt.Printf("accept retrying in %v", tempDelay)
	time.Sleep(tempDelay)
	this.tempDelay = tempDelay
}

func (this *server) onNewSession(conn net.Conn) {

	c := &session{conn: conn}
	this.Lock()
	this.sessionSet[c] = struct{}{}
	this.Unlock()

	this.wgSession.Add(1)
	go c.recvData(this.wgSession)

}

func (this *server) run() {
	this.wgLn.Add(1)
	defer this.wgLn.Done()
	for {
		conn, err := this.ln.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && !ne.Temporary() {
				fmt.Errorf("Accept err: %v\n")
				return
			}
			this.ProcTempErr()
			continue
		}
		this.onNewSession(conn)
	}
}

func (this *server) Listen() {
	ln, err := net.Listen("tcp4", "0.0.0.0:8089")
	if err != nil {
		fmt.Println("listen fail: ", err)
		return
	}
	this.ln = ln
	this.sessionSet = make(SessionSet)
	go this.run()
}

func (this *session) recvData(wg sync.WaitGroup) {
	defer wg.Done()
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

func main() {
	svr := &server{}
	svr.Listen()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c

	svr.Close()
}

func (this *session) StartSession() {

}
