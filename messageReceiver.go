package main

import (
	"github.com/nats-io/go-nats"
	"os"
	"io/ioutil"
	"time"
	"fmt"
)

func main() {
	xmlFile, _ := os.Open("response_marsh.xml")
	defer xmlFile.Close()
	byteXML, _ := ioutil.ReadAll(xmlFile)
	nc,_ := nats.Connect(nats.DefaultURL)
	defer nc.Close()
	t0:=time.Now()
	for i:=0;i<1001;i++ {
		nc.Publish("messageHandlers", byteXML)
	}
	t1:=time.Now()
	fmt.Printf("Time passed:%s\r\n",t1.Sub(t0))
	//nc.QueueSubscribe("foo","queue1", func(m *nats.Msg) {
	//	fmt.Printf("Received a message: %s\n", string(m.Data))
	//})

}
