package main

import (
	"os"
	"io/ioutil"
	"fmt"
	//"./smisInteract/smisInteract.go"
	"awesomeProject/smisInteract"
)

func main() {

	//sig,_ := rsa.SignPKCS1v15(rand.Reader, pri, crypto.SHA1, d)

	xmlFile, _ := os.Open("response_marsh.xml")
	defer xmlFile.Close()
	byteXML, _ := ioutil.ReadAll(xmlFile)
	encryptedText,signature,publicKey,_, _:= smisInteract.ProcessXML(byteXML)
	fmt.Println(encryptedText)
	respText:=`<ns2:TestResponse xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns2="http://basis-edu.ru/monitoring/schemas/node" xmlns:xenc="http://www.w3.org/2001/04/xmlenc#"></ns2:TestResponse>`
	xmlResp, _ :=smisInteract.ConstructResponse(respText,string(signature),publicKey)
	//fmt.Println(xmlResp)
	ioutil.WriteFile("soapEnvelopeMarshalTesting.xml",[]byte(xmlResp),0644)
}



