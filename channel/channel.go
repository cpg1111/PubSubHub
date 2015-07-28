package channel

import(
    "net"
)

type channel struct{
    id string
    lateJoiner bool
    messages []string
    subscribers []net.Conn
}

func(c *channel) AddSubscriber(sub net.Conn){
    c.subscribers[len(c.subscribers)] = sub
    if c.lateJoiner {
        buffer := make([]byte, len(c.messages))
        for i :=0; i < len(c.messages); i++ {
            buffer[i] = c.messages[i]
        }
        sub.Write(buffer)
    }
}

func(c *channel) Publish(message String){
    for i := 0; i < len(c.subscribers); i++ {
        c.subscribers[i].Write(message)
    }
}
