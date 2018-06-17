package smisInteract

import (
	"strings"
	"regexp"
)

func ExcC14N(xmlText []byte) ([]byte,error) {
	//var nameSpace string
	xmlnsFound := false
	//var addToEnd string
	var idx int
	var tokName []string
	var tokParts []string
	var parentNsStack StringStack
	//var tokNs string
	//var xmlnsAttr string
	nameSpacesMap := make(map[string]string)
	nameSpacesMap["wsse"]="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	nameSpacesMap["wsse11"]="http://docs.oasis-open.org/wss/oasis-wss-wssecurity-secext-1.1.xsd"
	nameSpacesMap["ds"]="http://www.w3.org/2000/09/xmldsig#"
	nameSpacesMap["xenc"]="http://www.w3.org/2001/04/xmlenc#"
	//nameSpacesMap["wsse"]="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	nameSpacesMap["SOAP-ENV"]="http://schemas.xmlsoap.org/soap/envelope/"
	xmlToCanonString := string(xmlText)

	var textToInsert string
	var tokNsCur string
	var attrNs string
	r2 := regexp.MustCompile(`<(.*?)?>`)
	res2 := r2.FindAll(xmlText,-1)
	for res2Idx,toks := range res2 {
		textToInsert=""
		idx = strings.Index(xmlToCanonString,string(toks))
		tokParts=strings.Split(xmlToCanonString[idx:idx+len(toks)]," ")
		for tokPartIdx,tokPart := range tokParts {
			if tokPartIdx == 0 {
				tokNsCur=strings.Split(tokPart,":")[0][1:]
			}
			if (strings.Contains(tokPart,">")) {
				tokPart = tokPart[:len(tokPart)-1]
			}
			if (strings.Contains(tokPart,"xmlns"))&&!(strings.Contains(tokPart,"wsu")) {
				attrNs=strings.Split(tokPart,"=")[0][6:]
				if (strings.Compare(tokNsCur,attrNs)==0) {
					xmlnsFound = true
					if (len(parentNsStack.data)==0) {
						textToInsert = " " + tokPart + textToInsert
						continue
					}
					if (tokNsCur!=parentNsStack.data[len(parentNsStack.data)-1]) {
						textToInsert = " " + tokPart + textToInsert
					}
				}
				// need to check that no spaces are before the "<" sign
			} else if !(strings.Contains(tokPart,"<")) {
				textToInsert += " " + tokPart
			}
		}
		if (!xmlnsFound) &&((res2Idx==0)||((tokNsCur!=parentNsStack.data[len(parentNsStack.data)-1]))&& (!strings.Contains(string(toks),"</"))){
			textToInsert = " " + "xmlns:"+tokNsCur+"=\""+nameSpacesMap[tokNsCur]+"\""+textToInsert
		}
		if len(tokParts)>1 {
			textToInsert=tokParts[0] + textToInsert
		} else {
			textToInsert=tokParts[0][:len(tokParts[0])-1] + textToInsert
		}
		if strings.Contains(string(toks),"</") {
			parentNsStack.Pop()
			continue
		}
		if !strings.Contains(string(toks),`/>`) {

			parentNsStack.Push(tokNsCur)
		} else {
			tokName=strings.Split(xmlToCanonString[idx:idx+len(toks)]," ")
			textToInsert = textToInsert[:len(textToInsert)-1]
			textToInsert += "></"+tokName[0][1:]
		}
		xmlToCanonString = xmlToCanonString[:idx] + textToInsert + ">" + xmlToCanonString[idx+len(toks):]
	}
	//idx = strings.Index(xmlToCanonString,string(res2[0]))
	//tokNs=strings.Split(xmlToCanonString[idx:idx+len(res2[0])],":")[0]
	//tokParts=strings.Split(xmlToCanonString[idx:idx+len(res2[0])]," ")
	//var tokAttrsAndEnd string
	//if len(tokParts)>1 {
	//	tokAttrsAndEnd = " "+strings.Join(tokParts[1:]," ")
	//	tokAttrsAndEnd = tokAttrsAndEnd[:len(tokAttrsAndEnd)-1]
	//} else {
	//	tokParts[0]=tokParts[0][:len(tokParts[0])-1]
	//}
	//nameSpace=nameSpacesMap[tokNs[1:]]
	//xmlnsAttr=" xmlns:"+tokNs[1:]+"=\""+nameSpace+"\""

	//r := regexp.MustCompile(`<(.*?)?/>`)
	//res := r.FindAll([]byte(xmlToCanonString),-1)
	//
	//for _,toks := range res {
	//	fmt.Println(string(toks))
	//	idx = strings.Index(xmlToCanonString,string(toks))
	//	tokName=strings.Split(xmlToCanonString[idx:idx+len(toks)]," ")
	//	if len(xmlToCanonString)>idx+len(toks) {
	//		addToEnd = xmlToCanonString[idx+len(toks)+1:]
	//	}
	//	xmlToCanonString= xmlToCanonString[:idx]+xmlToCanonString[idx:idx+len(toks)-2]+"></"+tokName[0][1:]+">"+addToEnd
	//}
	xmlToCanonString=strings.Replace(xmlToCanonString,"\r\n","\n",-1)

	return []byte(xmlToCanonString),nil
}

