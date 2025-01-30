package transport

import (
	"net"
	"sync"
	"time"
)

type ConnPool struct {
	address  string
	mu       sync.Mutex
	pool     chan net.Conn
	maxConns int
}

func NewConnPool(address string, maxCoons int) *ConnPool {
	return &ConnPool{
		address:  address,
		maxConns: maxCoons,
		pool:     make(chan net.Conn, maxCoons),
	}
}
func (cp *ConnPool) Get() (net.Conn, error) {
	select {
	case conn := <-cp.pool:
		return conn, nil
	default:
		return net.DialTimeout("tcp", cp.address, 2*time.Second)
	}
}
func (cp *ConnPool) Put(conn net.Conn) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	select {
	case cp.pool <- conn:
	default:
		err := conn.Close()
		if err != nil {
			return
		}
	}
}
func (cp *ConnPool) Close() {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	close(cp.pool)
	for conn := range cp.pool {
		err := conn.Close()
		if err != nil {
			return
		}
	}
}
