package main

import (
	"regexp"
	"strings"
)

func changeBodyToDecrypted(decryptedBody string, byteXML []byte) []byte {
	r := regexp.MustCompile(`</?SOAP-ENV:Body(.*?)?>`)
	res := r.FindAll(byteXML,-1)
	byteXMLString:=string(byteXML)
	idx := strings.Index(byteXMLString,string(res[0]))
	idx2 := strings.Index(byteXMLString,string(res[1]))
	return []byte(byteXMLString[idx:idx+len(res[0])]+decryptedBody+byteXMLString[idx2:idx2+len(res[1])])
}
