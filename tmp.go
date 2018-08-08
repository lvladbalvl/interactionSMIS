package main

import (
	"os"
	"io/ioutil"
)

func tmp() {
	xmlFile, _ := os.Open("soapEnvelopeMarshalTesting.xml")
	defer xmlFile.Close()
	byteXML, _ := ioutil.ReadAll(xmlFile)

	//xmlFileExpected, _ := os.Open("afterCanonSoapBody.xml")
	//defer xmlFileExpected.Close()
	//xmlToCanon,_ := ioutil.ReadAll(xmlFileBeforeCanon)
	//xmlToCanonString := string(xmlToCanon)
	//signToCanon := SignedInfo{}
	//xml.Unmarshal(xmlToCanon,&signToCanon)
	//exptectedXML, _ := ioutil.ReadAll(xmlFileExpected)
	//canonedXML,_ := ExcC14N(xmlToCanon)
	//fmt.Println("sdf")
	//fmt.Println(string(canonedXML))
	//fmt.Println(exptectedXML)
	//sign:=SignedInfoRet{}
	//sign.CanonicalizationMethodRet.Algorithm=signToCanon.CanonicalizationMethod.Algorithm
	//sign.SignatureMethodRet.Algorithm=signToCanon.SignatureMethod.Algorithm
	//sign.ReferenceRet.URI=signToCanon.Reference.URI
	//sign.ReferenceRet.ValueType=signToCanon.Reference.ValueType
	//sign.ReferenceRet.TransformsRet.TransformRet.Algorithm=signToCanon.Reference.Transforms.Transform.Algorithm
	//sign.ReferenceRet.DigestMethodRet.Algorithm=sign.ReferenceRet.DigestMethodRet.Algorithm
	//text,_:=xml.Marshal(sign)
	//fmt.Println(string(text))
	//fmt.Println(xmlToCanonString)
	//r := regexp.MustCompile(`<(.+)?/>`)
	//res := r.FindAll(xmlToCanon,-1)
	//nameSpacesMap := make(map[string]string)
	//nameSpacesMap["wsse"]="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	//nameSpacesMap["ds"]="http://www.w3.org/2000/09/xmldsig#"
	//nameSpacesMap["xenc"]="http://www.w3.org/2001/04/xmlenc#"
	////nameSpacesMap["wsse"]="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	//nameSpacesMap["SOAP-ENV"]="http://schemas.xmlsoap.org/soap/envelope/"
	//var idx int
	//var tokName []string
	//fmt.Println("sdf")
	//var tokNs string
	//var nameSpace string
	//var xmlnsAttr string
	//r2 := regexp.MustCompile(`<(.+)?>`)
	//res2 := r2.FindAll(xmlToCanon,1)
	//idx = strings.Index(xmlToCanonString,string(res2[0]))
	//tokNs=strings.Split(xmlToCanonString[idx:idx+len(res2[0])],":")[0]
	//nameSpace=nameSpacesMap[tokNs[1:]]
	//xmlnsAttr=" xmlns:"+tokNs[1:]+"=\""+nameSpace+"\">"
	//xmlToCanonString = xmlToCanonString[:idx+len(res2[0])-1]+xmlnsAttr+xmlToCanonString[idx+len(res2[0]):]
	//for _,toks := range res {
	//	idx = strings.Index(xmlToCanonString,string(toks))
	//	tokName=strings.Split(xmlToCanonString[idx:idx+len(toks)]," ")
	//	xmlToCanonString= xmlToCanonString[:idx]+xmlToCanonString[idx:idx+len(toks)-2]+"></"+tokName[0][1:]+">"+xmlToCanonString[idx+len(toks)+1:]
	//}
	//strings.Replace(xmlToCanonString,"\r\n","\n",-1)

}
