package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
	"bytes"
	"encoding/xml"
	"io/ioutil"
)

// need to define list of xml.Names that don't have xmlns attribute
// need to strip tabs in content
// need to delete and self-close tags from the other list
// need to delete empty objects

func xmlAdditionalConstructor(input string) string {
	var nsStack StringStack //stack to save namespaces to apply for closing tags
	var stringToWrite string //a variable used to replace /n to /r/n in data inside tags
	var addRet string //a variable to use when the /r/n is needed after a tag
	nameSpace :="" // xml marshals tokens without nameSpaces - need to add them to each one. getNamespace function gets a nameSpace by matching xmlns attribute
	toDeleteKeyIdent := false // there are 2 SecurityTokenReference in an answer. One has KeyIdentifier token and the other one has Reference attribute.
	//So these to bool vars are required to delete both start and end element of a token. The deletion is done if a token has only one attribute (xlmns)
	toDeleteReference := false
	securityTokenAddAttribSwitch := true // First SecurityTokenReference has attribute wsu:Id and the second one does not have.
	// So need to add this attributes only for the first time and make this value false
	// For other tokens this could be achieved my omitempty tag but this attribute (wsu:Id) is namespaced - can be read bu unmarshal but can not be marshalled properly
	listOfTokensWithoutNs := make([]string,18) //list of tokens which don't have an xmlns: attribute (marshalled by default by xml.Marshal). May be not needed - check later
	// two of tokens require namespace for the first time but don't require for the second!
	listOfTokensWithoutNs[0]="EncryptionMethod"
	listOfTokensWithoutNs[1]="KeyIdentifier"
	listOfTokensWithoutNs[2]="CipherData"
	listOfTokensWithoutNs[3]="CipherValue"
	listOfTokensWithoutNs[4]="ReferenceList"
	listOfTokensWithoutNs[5]="DataReference"
	listOfTokensWithoutNs[6] = "Header"
	listOfTokensWithoutNs[7] = "Body"
	listOfTokensWithoutNs[8] = "SignedInfo"
	listOfTokensWithoutNs[9] = "Transforms"
	listOfTokensWithoutNs[10] = "Transform"
	listOfTokensWithoutNs[11] = "SignatureValue"
	listOfTokensWithoutNs[12] = "CanonicalizationMethod"
	listOfTokensWithoutNs[13] = "SignatureMethod"
	listOfTokensWithoutNs[14] = "DigestMethod"
	listOfTokensWithoutNs[15] = "DigestValue"
	listOfSelfClosingTokens := make([]string,6) // list of self-closing tags. xml.Marshal makes all tokens to be with start and end elements - no self-closing ones.
	// they write on the forums that if it is no difference - why to use self-closings. May be not needed - check later
	listOfSelfClosingTokens[0] = "CanonicalizationMethod"
	listOfSelfClosingTokens[1] = "SignatureMethod"
	listOfSelfClosingTokens[2] = "Transform"
	listOfSelfClosingTokens[3] = "DigestMethod"
	listOfSelfClosingTokens[4] = "EncryptionMethod"
	listOfSelfClosingTokens[5] = "DataReference"
	listOfTokensWithoutFirstReturn := make([]string,4) // the main problem was with returns!!! They should match exactly or will cause an error
	// Here is the list of tokens which don't have return after first token, but have one after the last!
	listOfTokensWithoutFirstReturn[0] = "DigestValue"
	listOfTokensWithoutFirstReturn[1] = "SignatureValue"
	listOfTokensWithoutFirstReturn[2] = "SecurityTokenReference"
	listOfTokensWithoutLastReturn := make([]string,2) // list of tokens which have a return after start token!!! but don't have a return after last token!
	// BUT KeyInfo for the first time in not in this list but is added after the first occurance!!!
	listOfTokensWithoutLastReturn[0] = "Signature"
	file, _ := os.Create("testing.xml") // buffers in file. WIll change it later
	defer file.Close()
	writer := bufio.NewWriterSize(bufio.NewWriter(file),4096*4) // this magic is needed as bufio has only 4096 bytes reserved for buffering for default
	inputReader:=strings.NewReader(input) // now let's parse already marshalled soap envelope
	decoder := xml.NewDecoder(inputReader)
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		// Switch a type of a token (start, end element or data inside)
		switch x := t.(type) {
		case xml.StartElement:
			nameSpace = getNamespace(x.Attr) // first get namespace from map of xmlns attribute types
			addRet = "" // addRet (adding \r\n) - is empty by default
			if (nameSpace == "ds") && !stringInSlice(x.Name.Local,listOfTokensWithoutFirstReturn){ //most often tokens with ds namespace have a return
			// but there are some which don't have (they are in listOfTokensWithoutFirstReturn)
				addRet = "\r\n"
			}
			// two next checks are to delete either KeyIdentifier or Reference from SecurityTokenReference and to enable toDelete bool vars - for deletion of end tokens
			if (x.Name.Local == "KeyIdentifier") && (len(x.Attr)<2) {
				toDeleteKeyIdent = true
				continue;
			}
			if (x.Name.Local == "Reference") && (len(x.Attr)<2) {
				toDeleteReference = true
				continue;
			}
			// push nameSpace in a stack to use for the end token
			nsStack.Push(nameSpace)
			// print the beginning of the token
			fmt.Fprintf(writer,"<%s:%s", nameSpace,x.Name.Local)
			if x.Name.Local == "BinarySecurityToken" {
			}
			// iterate through attribute list and as Signature and KeyInfo have nameSpace only in the first occurence - add them to list
			// if token is in the listOfTokensWithoutNs - don't add xlmns attribute
			for _, attr := range x.Attr {
				if (attr.Name.Local == "xmlns") {
					if (stringInSlice(x.Name.Local,listOfTokensWithoutNs)) {
						if x.Name.Local == "Signature" {
							listOfTokensWithoutNs[16] = "Signature"
						} else if x.Name.Local == "KeyInfo" {
							listOfTokensWithoutNs[17] = "KeyInfo"
						}
						continue
					}
				}
				// add attribure name. for xmlns add also a namespace. xml Marshal writes it like xlmns="http:.... but we need xmlns:wsse="http:...
				fmt.Fprintf(writer," %s", attr.Name.Local)
				if (attr.Name.Local == "xmlns") {
					fmt.Fprint(writer,":"+nameSpace)
				}
				//start attribute description with ="
				writer.Write([]byte("=\""))
				// print attribute value
				xml.EscapeText(writer, []byte(attr.Value))
				// end with "
				writer.Write([]byte{'"'})
			}
			//switch the x.Name.Local - for different tokens add different attributes which are namespaced and can not be marshalled.
			// for SecurityTokenReference - only the first occurence has the attributes - make a bool variable securityTokenAddAttribSwitch false
			switch x.Name.Local {
			case "BinarySecurityToken":
				fmt.Fprint(writer," xmlns:wsu=\"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd\"")
				fmt.Fprint(writer," wsu:Id=\"CertId-63D3AB8DA51E36C344152767385959333533\"")
			case "SecurityTokenReference":
				if securityTokenAddAttribSwitch {
					securityTokenAddAttribSwitch = false
					fmt.Fprint(writer," xmlns:wsu=\"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd\"")
					fmt.Fprint(writer," wsu:Id=\"STRId-63D3AB8DA51E36C344152767385959333535\"")
				}
			case "Envelope":
				fmt.Fprint(writer," xmlns:xenc=\"http://www.w3.org/2001/04/xmlenc#\"")
			case "Body":
				fmt.Fprint(writer," xmlns:wsu=\"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd\"")
				fmt.Fprint(writer," wsu:Id=\"id-20121\"")
			case "Security":
				fmt.Fprint(writer," SOAP-ENV:mustUnderstand=\"1\"")
			}
			// end a start element with > and /> for self-closing ones
			if !stringInSlice(x.Name.Local,listOfSelfClosingTokens) && !(x.Name.Local == "Reference" && nameSpace == "wsse") {
				fmt.Fprintf(writer, ">"+addRet)
			} else {
				fmt.Fprintf(writer, "/>"+addRet)
		}
		case xml.CharData:
			//delete spaces and replace \n with \r\n
			stringToWrite=string(bytes.TrimSpace(x))
			if (len(stringToWrite)>0) {
				stringToWrite=(strings.Replace(string(x),"\n","\r\n",-1))
			}
			fmt.Fprintf(writer, stringToWrite)
		case xml.EndElement:
			// this is the place where we needed toDelete bool variables - now the end tokens are also deleted and these variable get false again
			if (x.Name.Local == "KeyIdentifier") && (toDeleteKeyIdent) {
				toDeleteKeyIdent=false
				continue;
			}
			if (x.Name.Local == "Reference") && (toDeleteReference) {
				toDeleteReference=false
				continue;
			}
			//get namespace for a closing token from a stack (deleting it)
			nameSpace = nsStack.Pop()
			// add \r\n if a token is in listOfTokensWithoutFirstReturn - these tokens don't have first return but do have last
			// also add \r\n for ds tokens
			if stringInSlice(x.Name.Local,listOfTokensWithoutFirstReturn) {
				addRet="\r\n"
			} else if (nameSpace == "ds") {
				addRet="\r\n"
			}
			// BUT if the token is in the listOfTokensWithoutLastReturn - no return even if it is a ds token
			if (stringInSlice(x.Name.Local,listOfTokensWithoutLastReturn)) {
				addRet=""
			}
			// if not a self-closing token - print it. but the problem is that Reference has two types: wsse:Reference and ds:Reference
			// so it can not be added to listOfSelfClosingTokens - because this would result in both of them being self-closing
			// while only wsse:Reference should - so additional check required
			if !stringInSlice(x.Name.Local,listOfSelfClosingTokens) && !(x.Name.Local == "Reference" && nameSpace == "wsse") {
				fmt.Fprintf(writer, "</%s:%s>"+addRet, nameSpace, x.Name.Local)
			}
			// first occurence of KeyInfo has a return after it and the other ones don't have. So add it to the listOfTokensWithoutLastReturn
			if (x.Name.Local == "KeyInfo") {
				listOfTokensWithoutLastReturn[1] = "KeyInfo"
			}
		}
	}
	// dumping all the data to the file. then reading it and outputing
	writer.Flush()
	resultingXML,_:=ioutil.ReadFile("testing.xml")
	return string(resultingXML)
}
func getNamespace(attrs []xml.Attr) string {
	nameSpacesMap := make(map[string]string)
	nameSpacesMap["http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"]="wsse"
	nameSpacesMap["http://www.w3.org/2000/09/xmldsig#"]="ds"
	nameSpacesMap["http://www.w3.org/2001/04/xmlenc#"]="xenc"
	nameSpacesMap["http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"]="wsse"
	nameSpacesMap["http://schemas.xmlsoap.org/soap/envelope/"]="SOAP-ENV"
	for _, attr := range attrs {
		if attr.Name.Local == "xmlns"{
			return nameSpacesMap[attr.Value]
		}
	}
	return ""
}
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}