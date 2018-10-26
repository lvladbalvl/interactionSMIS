package microServs

import (
	"github.com/nats-io/go-nats"
	"encoding/json"
	"../smisInteract"
	"runtime"
	"encoding/base64"
	"crypto/x509"
	"crypto/rsa"
	"fmt"
)

func MsgConstructor() {
	nc,_ := nats.Connect(nats.DefaultURL)
	nc.QueueSubscribe("messageConstructors","queueToConstructors", func(m *nats.Msg) {
		inMessage:=Message{}
		json.Unmarshal(m.Data,&inMessage)
		certBytes:=inMessage.PubKey
		certBlock := make([]byte, base64.StdEncoding.DecodedLen(len(certBytes)))
		n, _ := base64.StdEncoding.Decode(certBlock, certBytes)
		cert,_ := x509.ParseCertificate(certBlock[:n])
		//fmt.Println(cert.SignatureAlgorithm.String())
		//err =cert.CheckSignatureFrom(cert)
		//fmt.Println(err)
		rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)
		xmlResp, _ :=smisInteract.ConstructResponse(string(inMessage.Text),string(inMessage.Signature),rsaPublicKey)
		outMessage := Message{[]byte(xmlResp),[]byte(""),[]byte("")}
		jsonMsg,err :=json.Marshal(outMessage)
		if (err!=nil) {
			fmt.Println(err)
		}
		nc.Publish(m.Reply,jsonMsg)
	})
	runtime.Goexit()
}
