package connection

import (
	"github.com/cpg1111/PubSubHub/room"

	"golang.org/x/net/websocket"
)

type WSHandler struct {
	Connection
}

func (h *WSHandler) HandleConnection(conn websocket.Conn, rooms map[string]room.Room) {
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
					rooms[id] = room.New(id, false)
				case "room.message":
					room.Publish(incomingMessage)
				case "room.destroy":
					rooms[incomingMessage.Room] = nil
				}
			}
		}()
	}
}
