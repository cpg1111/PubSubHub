package connection

import (
	"encoding/json"
	"net"
    "strings"

	"github.com/cpg1111/PubSubHub/room"

	"code.google.com/p/go-uuid/uuid"
)

type TcpHandler struct {
	Connection
    Conn net.Conn
}

func(h *TcpHandler) NewRoom(isMaster bool, createMessage Incoming{}, rooms *map[string]room.Channel) room.Channel{
    isLJ := false
    if strings.Contains(createMessage.content, "LateJoiner: true") {
        isLJ = true
    }
    newRoom := room.New(uuid.New(), isLJ)
    rooms[newRoom.Id] = newRoom
    if !isMaster {

    }
    return newRoom
}

func(h *TcpHandler) HandleConnection(conn net.Conn, rooms *map[string]room.Channel) {
	decoder := json.NewDecoder(conn)
	incomingMessage := &Incoming{}
	err := decoder.Decode(incomingMessage)
	if err != nil {
		panic(err)
	}
	switch incomingMessage.MessType {
	case "room.create":
		id := uuid.New()
		rooms[id] = room.New(id, false) // TODO late joiner option on create
	case "room.message":
		rooms[incomingMessage.Room].Publish(incomingMessage)
	case "room.destroy":
		rooms[incomingMessage.Room] = nil
	}
}

func(h *TcpHandler) Serve(connStr string, rooms *map[string]room.Channel){
    listener, lErr := net.Listen("tcp", connStr)
    if lErr != nil {
        panic(lErr)
    }
    defer listener.Close()
    for {
        conn, cErr := listener.Accept()
        if cErr != nil {
            panic(cErr)
        }
        go h.HandleConnection(conn, rooms)
    }
}

func(h *TcpHandler) Connect(addr string){
    conn, cErr := net.Dial("tcp", addr)
    if cErr != nil {
        panic(cErr)
    }
    h.Conn = conn
}

func(h *TcpHandler) Send(messType, messContent, messRoom string){
    mess := &Incoming{messType, messRoom, messContent}
    encoder := json.NewEncoder(h.Conn)
    encoder.Encode(mess)
}

func NewTCP() *TcpHandler {
    return &TcpHandler{}
}
