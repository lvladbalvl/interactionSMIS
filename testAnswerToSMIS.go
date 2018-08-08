package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"encoding/xml"
)

func testanswerToSMIS() {
	xmlFile, _ := os.Open("responseBack.xml")
	soapEnvStructure := SoapEnvelope2{}
	defer xmlFile.Close()
	byteXML, _ := ioutil.ReadAll(xmlFile)
	if err := xml.Unmarshal(byteXML, &soapEnvStructure); err != nil {
		fmt.Println(err)
	}
	fmt.Println(soapEnvStructure.Header.Security.Signature.KeyInfo.SecurityTokenReference.X509Data)
	output, err := xml.MarshalIndent(soapEnvStructure, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	_ = ioutil.WriteFile("responseBack_marsh.xml", []byte(xmlAdditionalConstructor(string(output))), 0644)
}
