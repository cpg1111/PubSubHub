package main

import(
	"log"
	"net"
	"os"
	"io/ioutil"
	"path/filepath"
	"encoding/json"
	"code.google.com/p/go-uuid/uuid"
	"./configReader"
	"./channel"
)

func main() {
	PATH_TO_CONFIG, pathErr := filepath.abs("/etc/pubsubhub/config.yml")
	alt_conf_path = os.GetEnv("PSH_CONF_PATH")
	if alt_conf_path != nil {
		PATH_TO_CONFIG = filepath.abs(alt_conf_path)
	}
	if pathErr != nil {
		log.Fatal(pathErr)
	}
	conf := config{}
	conf.defaults()
	confFile, err := ioutil.ReadFile(PATH_TO_CONFIG)
	yamlErr := yaml.Unmarshal([]byte(data), &t)
	if yamlErr != nil {
		log.Fatal(yamlErr)
	}
	/*ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	channels := make(map[string]channel.Channel)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn, channels)
	}*/
}

type incoming struct{
	messType string
	channel string
	content string
}

func handleConnection(conn net.Conn, channels map[string]channel.Channel){
	tmpBuffer := make([]byte, 256)
	maxByteLength, err := conn.Read(tmpBuffer)
	buffer := make([]byte, 0, maxByteLength)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	for {
		message, err := conn.Read(buffer)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		incomingMessage := incoming{}
		parsedMessage := json.Unmarshal(message, &incomingMessage)
		switch parseMessage.messType {
			case "create_channel":
				newUuid := uuid.New()
				shouldHaveLateJoiner = parseMessage.content == "later_joiner"
				channels[newUuid] = channel.Channel{Id: newUuid, LateJoiner: ShouldHaveLateJoiner, Messages: make([]string), Subscribers: make([conn]net.Conn)}
			case "publish":
				channels[parseMessage.channel].Publish(content)
			case "subscribe":
				channels[parseMessage.channel].AddSubscriber(conn)
		}
	}
}
