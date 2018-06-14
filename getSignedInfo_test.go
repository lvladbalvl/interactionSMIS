package main

import (
	"testing"
	"os"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

func TestGetSignedInfo(t *testing.T) {
	asserter := assert.New(t)
	xmlFile, _ := os.Open("response (2).xml")
	defer xmlFile.Close()
	byteXML, _ := ioutil.ReadAll(xmlFile)
	signedInfo,_ := getSignedInfo(byteXML)
	xmlFileExpected, _ := os.Open("beforeCanonSignedInfo.xml")
	defer xmlFileExpected.Close()
	expectedXML, _ := ioutil.ReadAll(xmlFileExpected)
	asserter.Equal( string(expectedXML), signedInfo, "The two xml docs should be the same.")

}
