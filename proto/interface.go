package proto

import (
	"github.com/zk-123/ngrok/conn"
)

type Protocol interface {
	GetName() string
	WrapConn(conn.Conn, interface{}) conn.Conn
}
