package main

import (
	"github.com/nats-io/go-nats"
	"awesomeProject/smisInteract"
	"encoding/json"
	"fmt"
	"runtime"
)
type Message struct {
	Text []byte `json:"text"`
	Signature []byte `json:"signature"`
	PubKey []byte	`json:"pubKey"`
}
func main() {
	nc,_ := nats.Connect(nats.DefaultURL)
	nc.QueueSubscribe("messageHandlers","queueToHandlers", func(m *nats.Msg) {
		encryptedText,signature,publicKey,_, _:= smisInteract.ProcessXML(m.Data)
		forwardMessage:=Message{[]byte(encryptedText),signature,publicKey}
		jsonMsg,err :=json.Marshal(forwardMessage)
		if err!=nil {
			fmt.Println(err)
		}
		nc.Publish("dbWorkers",jsonMsg)
	})
	runtime.Goexit()
}
