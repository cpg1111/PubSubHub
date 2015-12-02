package connection

import (
	"encoding/json"
	"net"

	"github.com/cpg1111/PubSubHub/room"

	"code.google.com/p/go-uuid/uuid"
)

type TcpHandler struct {
	Connection
}

func (h *TcpHandler) HandleConnection(conn net.Conn, rooms map[string]room.Room) {
	decoder := json.NewDecoder(conn)
	for {
		go func() {
			incomingMessage := &Incoming{}
			for {
				err := decoder.Decode(incomingMessage)
				if err != nil {
					panic(err)
				}
				switch incomingMessage.MessType {
				case "room.create":
					id := uuid.New()
					rooms[id] = room.New(id, false) // TODO late joiner option on create
				case "room.message":
					room.Publish(incomingMessage)
				case "room.destroy":
					rooms[incomingMessage.Room] = nil
				}
			}
		}()
	}
}
