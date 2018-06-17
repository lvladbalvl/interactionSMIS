package smisInteract

import "regexp"

func getSoapBody(byteXML []byte) ([]byte,error) {
	r := regexp.MustCompile(`<SOAP-ENV:Body((.|\r\n)*?)</SOAP-ENV:Body>`)
	res := r.FindAll(byteXML,-1)
	if len(res)==0 {
		return []byte(""),nil
	}
	//byteXMLString:=string(byteXML)
	//idx := strings.Index(byteXMLString,string(res[0]))
	return res[0],nil
}
