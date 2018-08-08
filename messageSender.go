package main

import (
	"github.com/nats-io/go-nats"
	"io/ioutil"
	"fmt"
	"runtime"
	"time"
)


func main() {
	k:=0
	nc,_ := nats.Connect(nats.DefaultURL)
	t0:=time.Now()
	nc.Subscribe("messageSender", func(m *nats.Msg) {
		ioutil.WriteFile("messageToSend.xml",m.Data,0644)
		k++
		if k%1000 ==1 {
			t0=time.Now()
		}
		if k%1000 ==0 {
			t1:=time.Now()
			fmt.Printf("Time passed:%s\r\n",t1.Sub(t0))

		}
		fmt.Printf("The number of sended messages is: %d\r\n",k)
	})
	runtime.Goexit()
}