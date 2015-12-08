package connection

import (
	"net"

    "github.com/cpg1111/PubSubHub/room"
)

type Incoming struct {
	MessType string
	Room     string
	Content  string
}

type Connection interface {
    NewRoom(isMaster bool, createMessage Incoming{}) room.Channel
	HandleConnection(conn *interface{}, rooms *map[string]room.Channel)
}

func New(connType, target string) *interface{}, *Connection {
    var conn *Connection
    switch connType {
    case "TCP":
        conn = &TcpHandler{}
        l, lErr := net.Listen("tcp", target)
    }
}
