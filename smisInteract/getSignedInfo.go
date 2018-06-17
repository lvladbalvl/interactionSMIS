package smisInteract

import (
	"regexp"
	"fmt"
)

func getSignedInfo(byteXML []byte) (string,error) {
	r := regexp.MustCompile(`<ds:SignedInfo>((.|\r\n)*?)</ds:SignedInfo>`)
	res := r.FindAll(byteXML,-1)
	if len(res)==0 {
		return "",fmt.Errorf("SOAP Body not found")
	}
	return string(res[0]),nil
}
