package microServs

import (
	"github.com/nats-io/go-nats"
	"../smisInteract"
	"encoding/json"
	"fmt"
	"runtime"
)
func MsgHandler() {
	nc,_ := nats.Connect(nats.DefaultURL)
	nc.QueueSubscribe("messageHandlers","queueToHandlers", func(m *nats.Msg) {
		inMessage:=Message{}
		json.Unmarshal(m.Data,&inMessage)
		encryptedText,signature,publicKey,_, _:= smisInteract.ProcessXML(inMessage.Text)
		forwardMessage:=Message{[]byte(encryptedText),signature,publicKey}
		jsonMsg,err :=json.Marshal(forwardMessage)
		if err!=nil {
			fmt.Println(err)
		}
		nc.PublishRequest("dbWorkers",m.Reply,jsonMsg)
	})
	runtime.Goexit()
}
