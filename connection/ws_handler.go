package connection;

import (
    "golang.org/x/net/websocket"
)

type WSHandler struct {
    Connection
}

func(h *WSHandler) HandleConnection(conn websocket.Conn, rooms map[string]room.Room) {

}
