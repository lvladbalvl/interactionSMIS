package smisInteract

import (
	"regexp"
	"fmt"
)

func getSoapBody(byteXML []byte) ([]byte,error) {
	r := regexp.MustCompile(`<SOAP-ENV:Body((.|\r\n)*?)</SOAP-ENV:Body>`)
	res := r.FindAll(byteXML,-1)
	if len(res)==0 {
		return []byte(""),fmt.Errorf("SOAP Body not found")
	}
	return res[0],nil
}
