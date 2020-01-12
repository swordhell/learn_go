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

type (
	// 会话对象
	session struct {
		sync.Mutex
		conn      net.Conn
		writeChan chan []byte
		closeFlag bool
		owner     *server
	}

	SessionSet map[*session]struct{}

	// 服务器对象
	server struct {
		sync.Mutex
		tempDelay  time.Duration
		ln         net.Listener
		wgLn       sync.WaitGroup
		mutexConns sync.Mutex
		conns      SessionSet
		wgSession  sync.WaitGroup
	}
)

func (this *server) Close() {
	fmt.Printf("server.Close() address: %v\n", this.ln.Addr().String())
	this.ln.Close()
	this.wgLn.Wait()

	this.mutexConns.Lock()
	for conn := range this.conns {
		conn.Close()
	}
	this.conns = nil
	this.mutexConns.Unlock()
	this.wgSession.Wait()
	fmt.Printf("server.Close() done\n", this.ln.Addr().String())
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
	fmt.Printf("server.ProcTempErr accept retrying in %v\n", tempDelay)
	time.Sleep(tempDelay)
	this.tempDelay = tempDelay

}

func (this *server) onNewSession(conn net.Conn) {
	fmt.Printf("server.onNewSession remote %v\n", conn.RemoteAddr().String())

	c := &session{conn: conn, closeFlag: false, owner: this}
	c.writeChan = make(chan []byte, 512*1024)

	this.mutexConns.Lock()
	this.conns[c] = struct{}{}
	this.mutexConns.Unlock()

	go c.goRecvData(this.wgSession)
	go c.goSendData(this.wgSession)
	fmt.Println("server.onNewSession done")

	c.PostData(`I had seen little of Holmes lately.
My marriage had drifted us away from each other.`)
}

func (this *server) onSessionLost(conn *session) {
	fmt.Println("server.onSessionLost")
	this.mutexConns.Lock()
	if _, ok := this.conns[conn]; ok {
		delete(this.conns, conn)
		fmt.Println("server.onSessionLost done")
	}
	this.mutexConns.Unlock()
}

func (this *server) runAccept() {
	this.wgLn.Add(1)
	defer this.wgLn.Done()
	for {
		conn, err := this.ln.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && !ne.Temporary() {
				fmt.Printf("Accept err: %v\n", ne.Error())
				return
			}
			this.ProcTempErr()
			continue
		}
		this.onNewSession(conn)
	}
	fmt.Println("server.runAccept done")
}

func (this *server) Listen() {
	ln, err := net.Listen("tcp4", "0.0.0.0:8089")
	if err != nil {
		fmt.Println("listen fail: ", err)
		return
	}
	this.ln = ln
	this.conns = make(SessionSet)
	go this.runAccept()
	fmt.Printf("server.Listen address: %v done\n", ln.Addr().String())
}

func (this *session) PostData(data string) {
	msgLen := len(data)
	buf := []byte(data)
	msg := make([]byte, 2+msgLen)
	binary.LittleEndian.PutUint16(msg, uint16(msgLen))
	copy(msg[2:], buf)

	this.Lock()
	if !this.closeFlag {
		this.writeChan <- msg
	}
	this.Unlock()
	fmt.Printf("session.PostData address: %v, data size: %v\n",
		this.conn.RemoteAddr().String(), len(msg))
}

func (this *session) Close() {
	fmt.Println("session.Close")
	this.innorCloseSock()
}

func (this *session) innorCloseSock() {
	this.Lock()
	if !this.closeFlag {
		this.closeFlag = true
		this.conn.Close()
		close(this.writeChan)
		fmt.Println("session innorCloseSock done")
	}
	this.Unlock()
}

func (this *session) goRecvData(wg sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	conn := this.conn
	for {
		data, err := procData(conn)
		if err != nil {
			fmt.Printf("session.goRecvData procData fail, err: %v\n", err)
			break
		}
		this.handleData(data)
	}
	this.innorCloseSock()
}

func (this *session) goSendData(wg sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	conn := this.conn

	for b := range this.writeChan {
		if b == nil {
			fmt.Println("session.goSendData b is nil")
			break
		}
		_, err := conn.Write(b)
		if err != nil {
			fmt.Printf("session.goSendData Write fail, err: %v\n", err)
			break
		}
	}
	this.owner.onSessionLost(this)
	this.innorCloseSock()
}

func procData(conn net.Conn) ([]byte, error) {
	var bufHeader []byte = make([]byte, 2, 2)
	if _, err := io.ReadFull(conn, bufHeader); err != nil {
		fmt.Printf("procData io.ReadFull header fail, err: %v\n", err.Error())
		return nil, err
	}
	var msgLen uint32
	msgLen = uint32(binary.LittleEndian.Uint16(bufHeader))
	fmt.Printf("procData header get size: %v\n", msgLen)
	msgData := make([]byte, msgLen)
	if _, err := io.ReadFull(conn, msgData); err != nil {
		fmt.Printf("procData io.ReadFull body fail, err: %v\n", err.Error())
		return nil, err
	}
	fmt.Printf("procData get data, size: %v\n", len(msgData))
	return msgData, nil
}

func (this *session) handleData(data []byte) {
	fmt.Printf("session.handleData remote address: %v, recv data size: %v\n",
		this.conn.RemoteAddr().String(), len(data))
	content := string(data)
	fmt.Println(content)
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
