package main

import(
	"fmt"
	"os"
	"reflect"
	"net"
	"json"
	"code.google.com/p/go-uuid/uuid"
	"channel"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	channels := map[string]channel
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn, channels)
	}
}

type incoming interface{
	messType string
	channel string
	content string
}

func handleConnection(conn Conn, channels map[string]channel){
	tmpBuffer := make([]byte, 256)
	maxByteLength, err := conn.Read(tmpBuffer)
	buffer := make([]byte, 0, maxByteLength)
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		message, err := conn.Read(buffer)
		if err != nil {
			log.Fatal(err)
			return
		}
		parsedMessage := json.Unmarshall(message, incoming)
		switch parseMessage.messType {
			case "create_channel":
				newUuid := uuid.New()
				shouldHaveLateJoiner = parseMessage.content == "later_joiner"
				channels[newUuid] = channel{id: newUuid, lateJoiner: shouldHaveLateJoiner, messages: []string, subscribers: [conn]net.Conn}
			case "publish":
				channels[]
		}
	}
}
