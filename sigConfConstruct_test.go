package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"io/ioutil"
	"encoding/xml"
)

func TestConstructSigConf(t *testing.T) {
	asserter := assert.New(t)
	signature := "HWh1wxmYNBVhEHoskC65BmKHYhq2XqiSkkSq6I4XTH9fs2D01pFIP/SqKnApAwI/JyoOlcAa18KJs9lCD7BVWcSwqQJ3NpTxfayaeO8Z9TzcHkkXjhqH0IZMYwKtondF9AjoRTNclrM5yLKhK9NHnoUAH96mT5FwhEbGlyXyusNEO9CLs9HKXZoY+jPHiNFSIVDo1lh3+QJszO97OLaHq1vZXojxMNq1PVcAzWDQGYg/RR9D2zCCQWtQH6PD6ZQebrJ7dZfzf1AMAJ0scbAx9YeRpyCI6kkdP3EFXKKP2npVBHokHqLfHkxz1mQoFbLq+hML1ZA2ez6E+1svT5YJXg=="
	sigConf := SignatureConfirmation{}
	sigConf.Value = signature
	sigConf.WsuNs = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	sigConf.Xmlns = "http://docs.oasis-open.org/wss/oasis-wss-wssecurity-secext-1.1.xsd"
	sigConf.WsuId = "SigConf-40121"
	sigConfString,_ := xml.Marshal(sigConf)
	xmlFileExpected, _ := os.Open("sigConfSample.xml")
	defer xmlFileExpected.Close()
	expectedXML, _ := ioutil.ReadAll(xmlFileExpected)
	asserter.Equal( string(expectedXML), string(sigConfString), "The two xml docs should be the same.")

}
