package connection;

import(
    "net"
    "github.com/gorilla/websocket"
)

type connection interface {
    handleConnection(conn net.Conn, channels map[string]channel.Channel)
}

type wsConnection struct {
    socket
}
