package main

import (
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
)

func tcpClient() {
	xmlToSend := []byte(constructSoapAnswer())
	//_ = ioutil.WriteFile("response_marsh.html", xmlToSend, 0644)
	_ =ioutil.WriteFile("response_marsh.xml",xmlToSend,0644)
	xmlToCheck,_ := ioutil.ReadFile("response.xml")
	var i int
	for {
		if xmlToCheck[i]==xmlToSend[i]{
			i++
			continue
		}
		fmt.Println(string(xmlToCheck[i-20:i+20]))
		break
	}
	return
	//xmlToCheck,_:=ioutil.ReadFile("response.xml")

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://192.168.0.11:8080/monitoring/node/dispatch/", bytes.NewBuffer(xmlToSend))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Accept","text/xml, text/html, image/gif, image/jpeg, *; q=.2, */*; q=.2")
	req.Header.Add("User-Agent","Java/1.7.0_80")
	req.Header.Add("Soapaction","\"http://basis-edu.ru/monitoring/schemas/node/TestRequest\"")
	req.Header.Add("Host","192.168.0.8:8080/monitoring/node/dispatch/")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	_ = ioutil.WriteFile("response.html", body, 0644)
}