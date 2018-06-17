package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"os"
	"io/ioutil"
)

func TestXMLsignedInfoCanon(t *testing.T) {
	asserter := assert.New(t)
	xmlFileBeforeCanon, _ := os.Open("beforeCanonSignedInfo.xml")
	defer xmlFileBeforeCanon.Close()
	xmlFileExpected, _ := os.Open("afterCanonSignedInfo.xml")
	defer xmlFileExpected.Close()
	xmlToCanon,_ := ioutil.ReadAll(xmlFileBeforeCanon)
	exptectedXML, _ := ioutil.ReadAll(xmlFileExpected)
	canonedXML,_ := ExcC14N(xmlToCanon)
	asserter.Equal( string(exptectedXML),string(canonedXML), "The two xml docs should be the same.")
}
func TestXMLBodyCanon(t *testing.T) {
	asserter := assert.New(t)
	xmlFileBeforeCanon, _ := os.Open("beforeCanonSoapBody.xml")
	defer xmlFileBeforeCanon.Close()
	xmlFileExpected, _ := os.Open("afterCanonSoapBody.xml")
	defer xmlFileExpected.Close()
	xmlToCanon,_ := ioutil.ReadAll(xmlFileBeforeCanon)
	exptectedXML, _ := ioutil.ReadAll(xmlFileExpected)
	canonedXML,_ := ExcC14N(xmlToCanon)
	asserter.Equal( string(exptectedXML), string(canonedXML), "The two xml docs should be the same.")

}
func TestXMLSigConf(t *testing.T) {
	asserter := assert.New(t)
	xmlFileBeforeCanon, _ := os.Open("beforeCanonSignConf.xml")
	defer xmlFileBeforeCanon.Close()
	xmlFileExpected, _ := os.Open("afterCanonSignConf.xml")
	defer xmlFileExpected.Close()
	xmlToCanon,_ := ioutil.ReadAll(xmlFileBeforeCanon)
	exptectedXML, _ := ioutil.ReadAll(xmlFileExpected)
	canonedXML,_ := ExcC14N(xmlToCanon)
	asserter.Equal( string(exptectedXML), string(canonedXML), "The two xml docs should be the same.")

}
