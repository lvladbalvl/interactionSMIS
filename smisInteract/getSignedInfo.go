package smisInteract

import (
	"regexp"
)

func getSignedInfo(byteXML []byte) (string,error) {
	r := regexp.MustCompile(`<ds:SignedInfo>((.|\r\n)*?)</ds:SignedInfo>`)
	res := r.FindAll(byteXML,-1)
	if len(res)==0 {
		return "",nil
	}
	//byteXMLString:=string(byteXML)
	//idx := strings.Index(byteXMLString,string(res[0]))
	return string(res[0]),nil
}
