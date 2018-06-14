package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func main() {

	//sig,_ := rsa.SignPKCS1v15(rand.Reader, pri, crypto.SHA1, d)

	xmlFile, _ := os.Open("response_marsh.xml")
	defer xmlFile.Close()
	byteXML, _ := ioutil.ReadAll(xmlFile)
	encryptedText,signature,_, _, _:= processXML(byteXML)
	fmt.Println(encryptedText)
	respText:=`<ns2:TestResponse xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns2="http://basis-edu.ru/monitoring/schemas/node" xmlns:xenc="http://www.w3.org/2001/04/xmlenc#"></ns2:TestResponse>`
	xmlResp, _ :=constructResponse(respText,string(signature))
	//fmt.Println(xmlResp)
	ioutil.WriteFile("soapEnvelopeMarshalTesting.xml",[]byte(xmlResp),0644)
}



