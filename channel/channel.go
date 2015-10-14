package channel

import(
    "net"
)

type Channel struct{
    Id string
    LateJoiner bool
    Messages []string
    Subscribers []net.Conn
}

func(c *Channel) AddSubscriber(sub net.Conn){
    c.Subscribers[len(c.Subscribers)] = sub
    if c.LateJoiner {
        buffer := make([]byte, len(c.Messages))
        for i := 0; i < len(c.Messages); i++ {
            byteMessage := []byte(c.Messages[i])
            for j := 0; j < len(byteMessage); j++ {
                buffer[j] = byteMessage[j]
            }
            sub.Write(buffer)
        }
    }
}

func(c *Channel) Publish(message string){
    buffer := make([]byte, len([]byte(message)))
    byteMessage := []byte(message)
    for a := 0; a < len(buffer); a++ {
        buffer[a] = byteMessage[a]
    }
    for i := 0; i < len(c.Subscribers); i++ {
        c.Subscribers[i].Write(buffer)
    }
    if c.LateJoiner {
        c.Messages[len(c.Messages)] = message
    }
}
