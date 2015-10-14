package connection;

import (
    "net"
    "encoding/json"
    "code.google.com/p/go-uuid/uuid"
    "room"
)

type TcpHandler struct {
    Connection
}

func(h *TcpHandler) HandleConnection(conn net.Conn, rooms map[string]room.Room){
    decoder := json.NewDecoder(conn)
    for {
        go func(){
                incomingMessage := &Incoming{}
                for {
                    err := decoder.Decode(incomingMessage)
                    if err != nil {
                        panic(err)
                    }
                    switch incomingMessage.MessType {
                    case 'room.create':
                        id := uuid.New()
                        rooms[id] := room.New()
                    case 'room.message':
                        room.Publish(incomingMessage)
                    case 'room.destroy':
                        rooms[incomingMessage.Room] = nil
                    }
                }
            }()
    }
}
