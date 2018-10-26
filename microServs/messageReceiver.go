package microServs

import (
	"github.com/nats-io/go-nats"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"log"
	"fmt"
	"time"
)

func MsgReceiver() {
	http.HandleFunc("/", handlerCustom)
	log.Printf("start")
	err := http.ListenAndServe(":8085", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }

}



func handlerCustom(w http.ResponseWriter, r *http.Request) {
	
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	nc,_ := nats.Connect(nats.DefaultURL)
	defer nc.Close()
	forwardMessage:=Message{body,[]byte(""),[]byte("")}
	jsonMsg,_ :=json.Marshal(forwardMessage)
	reply, err := nc.Request("messageHandlers", jsonMsg,10*time.Second)
	inMessage:=Message{}
	json.Unmarshal(reply.Data,&inMessage)
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
	} else {
		//ioutil.WriteFile("resp_tmp.xml",inMessage.Text,666)
		fmt.Println(w.Header().Get("Content-Type"))
		w.Write(inMessage.Text)
	}


}