package connection

import (
	"github.com/gorilla/websocket"
	"net"
)

type Incoming struct {
	MessType string
	Room     string
	Content  string
}

type Connection interface {
	HandleConnection(conn interface{}, rooms map[string]room.Room)
}
