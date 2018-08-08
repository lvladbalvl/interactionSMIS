package main

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"awesomeProject/smisInteract"
	"github.com/pkg/profile"
)



func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	http.HandleFunc("/", handlerCustom)
	// Start the HTTPS server in a goroutine
	//go http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil)
	// Start the HTTP server
	//log.Fatal(http.ListenAndServe(":8080", nil))
	log.Printf("start")
	err := http.ListenAndServe(":8090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func handlerCustom(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	decryptedMessage,signatureText,publicKey,smisErrorMessage, smisError := smisInteract.ProcessXML(body)
	if smisError != nil {
		fmt.Println(smisErrorMessage)
	}
	fmt.Println(string(decryptedMessage))
	respText:=`<ns2:TestResponse xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns2="http://basis-edu.ru/monitoring/schemas/node" xmlns:xenc="http://www.w3.org/2001/04/xmlenc#"></ns2:TestResponse>`
	xmlResp, _ :=smisInteract.ConstructResponse(respText,string(signatureText),publicKey)
	w.Header().Add("Content-Type", "text/xml; charset=utf-8")
	w.Header().Add("Accept","text/xml, text/html, image/gif, image/jpeg, *; q=.2, */*; q=.2")
	w.Header().Add("User-Agent","Java/1.7.0_80")
	w.Header().Add("Soapaction","\"http://basis-edu.ru/monitoring/schemas/node/TestResponse\"")
	ioutil.WriteFile("testResponseToSMIS.xml",[]byte(xmlResp),0644)
	fmt.Fprintf(w,xmlResp)
	defer r.Body.Close()
}
