package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"encoding/xml"
)

func constructSoapAnswer() string {
	xmlFile, _ := os.Open("response.xml")
	soapEnvStructure := SoapEnvelope{}
	defer xmlFile.Close()
	byteXML, _ := ioutil.ReadAll(xmlFile)
	if err := xml.Unmarshal(byteXML, &soapEnvStructure); err != nil {
		fmt.Println(err)
	}
	output, err := xml.MarshalIndent(soapEnvStructure, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return xmlAdditionalConstructor(string(output))
}
