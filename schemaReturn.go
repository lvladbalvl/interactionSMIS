package main

import "encoding/xml"

type SignedInfoRet struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# ds:SignedInfo"`
	CanonicalizationMethodRet CanonicalizationMethodRet
	SignatureMethodRet SignatureMethodRet
	ReferenceRet ReferenceRet
}
type CanonicalizationMethodRet struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# ds:CanonicalizationMethod"`
	Algorithm string `xml:"Algorithm,attr,omitempty"`
}
type SignatureMethodRet struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# ds:SignatureMethod"`
	Algorithm string `xml:"Algorithm,attr,omitempty"`
}
type ReferenceRet struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# ds:Reference"`
	URI string `xml:"URI,attr,omitempty"`
	ValueType string `xml:"ValueType,attr,omitempty"`
	TransformsRet TransformsRet
	DigestMethodRet DigestMethodRet
	DigestValueRet DigestValueRet
}
type TransformsRet struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# ds:Transforms"`
	TransformRet TransformRet
}
type DigestMethodRet struct{
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# ds:DigestMethod"`
	Algorithm string `xml:"Algorithm,attr,omitempty"`
}
type DigestValueRet struct{
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# ds:DigestValue"`
	Algorithm string `xml:"Algorithm,attr,omitempty"`
	Contents []byte `xml:",innerxml"`
}
type TransformRet struct{
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# ds:Transform"`
	Algorithm string `xml:"Algorithm,attr,omitempty"`
}