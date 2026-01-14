package netmux

import (
	"context"
	"net"
	"sync/atomic"
)

type Muxer interface {
	net.Listener

	// Open 打开一个双向子流。
	Open(context.Context) (net.Conn, error)

	// RemoteAddr 远端节点地址。
	RemoteAddr() net.Addr

	IsClosed() bool

	// Protocol 返回通信协议类型，一般用于调试。
	//	- protocol: 标准的底层通信协议，如：tcp udp
	Protocol() string

	// Traffic 数据传输字节数。
	Traffic() (rx, tx uint64)
}

type trafficStat struct {
	rx, tx atomic.Uint64
}

func (t *trafficStat) incrRX(n int) {
	if n > 0 {
		t.rx.Add(uint64(n))
	}
}

func (t *trafficStat) incrTX(n int) {
	if n > 0 {
		t.tx.Add(uint64(n))
	}
}

func (t *trafficStat) load() (uint64, uint64) {
	return t.rx.Load(), t.tx.Load()
}
