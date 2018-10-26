package microServs

import (
	"github.com/nats-io/go-nats"
	"encoding/json"
	"fmt"
	"runtime"
	"../smisInteract"
)



func DbWorker() {
	nc,_ := nats.Connect(nats.DefaultURL)
	nc.QueueSubscribe("dbWorkers","queueToWorkers", func(m *nats.Msg) {
		inMessage:=Message{}
		json.Unmarshal(m.Data,&inMessage)
		//respText:=[]byte(`<ns2:TestResponse xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns2="http://basis-edu.ru/monitoring/schemas/node" xmlns:xenc="http://www.w3.org/2001/04/xmlenc#"></ns2:TestResponse>`)
		respText:=smisInteract.WriteToDB(inMessage.Text)
		outMessage := Message{[]byte(respText),inMessage.Signature,inMessage.PubKey}
		jsonMsg,err :=json.Marshal(outMessage)
		if (err!=nil) {
			fmt.Println(err)
		}
		nc.PublishRequest("messageConstructors",m.Reply,jsonMsg)
	})
	runtime.Goexit()
}
