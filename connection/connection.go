package connection;

import(
    "net"
    "github.com/gorilla/websocket"
)

type Incoming struct{
    MessType string
    Room string
    Content string
}

type Connection interface {
    HandleConnection(conn interface{}, rooms map[string]room.Room)
}
