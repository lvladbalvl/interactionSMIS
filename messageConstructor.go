package main

import (
	"github.com/nats-io/go-nats"
	"encoding/json"
	"awesomeProject/smisInteract"
	"runtime"
	"encoding/base64"
	"crypto/x509"
	"crypto/rsa"
)

type Message struct {
	Text []byte `json:"text"`
	Signature []byte `json:"signature"`
	PubKey []byte	`json:"pubKey"`
}
func main() {
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
		nc.Publish("messageSender",[]byte(xmlResp))
	})
	runtime.Goexit()
}
