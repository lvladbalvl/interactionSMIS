package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"io/ioutil"
)

func TestChangeBodyToDecrypted(t *testing.T) {
	xmlFile, _ := os.Open("response (2).xml")
	defer xmlFile.Close()
	byteXML, _ := ioutil.ReadAll(xmlFile)
	asserter := assert.New(t)
	xmlFileBeforeAdd, _ := os.Open("TestRequest.xml")
	defer xmlFileBeforeAdd.Close()
	decryptedBody, _ := ioutil.ReadAll(xmlFileBeforeAdd)
	xmlFileExpected, _ := os.Open("beforeCanonSoapBody.xml")
	defer xmlFileExpected.Close()
	exptectedXML, _ := ioutil.ReadAll(xmlFileExpected)
	SoapBodyWithWsuAttrs := changeBodyToDecrypted(string(decryptedBody),byteXML)
	asserter.Equal( string(exptectedXML), string(SoapBodyWithWsuAttrs), "The two xml docs should be the same.")

}

