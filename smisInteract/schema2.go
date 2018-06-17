package smisInteract

import "encoding/xml"

// soap-envelope might require xmlns:xenc attribute
type SoapEnvelope2 struct {
	XMLName xml.Name `xml:"SOAP-ENV:Envelope"`
	Xmlns string `xml:"xmlns:SOAP-ENV,attr,omitempty"`
	XmlnsXenc string `xml:"xmlns:xenc,attr,omitempty"`
	Header  SoapHeader2
	Body    SoapBody2
}
type SoapHeader2 struct {
	XMLName  xml.Name `xml:"SOAP-ENV:Header"`
	Security wsSecurity2
}
type SoapBody2 struct {
	XMLName  xml.Name `xml:"SOAP-ENV:Body"`
	WsuNs string `xml:"xmlns:wsu,attr"`
	WsuId string `xml:"wsu:Id,attr"`
	Contents []byte `xml:",innerxml"`
}
type wsSecurity2 struct {
	XMLName             xml.Name `xml:"wsse:Security"`
	Xmlns string `xml:"xmlns:wsse,attr"`
	MustUnderstand string `xml:"SOAP-ENV:mustUnderstand,attr,omitempty"`
	EncryptedKey        EncryptedKey2
	Signature			Signature2
	SignatureConfirmation SignatureConfirmation
}
type EncryptedKey2 struct {
	XMLName xml.Name `xml:"xenc:EncryptedKey"`
	Id string `xml:"Id,attr,omitempty"`
	Xmlns string `xml:"xmlns:xenc,attr"`
	EncryptionMethod EncryptionMethod2
	KeyInfo KeyInfo2
	CipherData CipherData2
	ReferenceList ReferenceList2
}
type Signature2 struct {
	XMLName xml.Name `xml:"ds:Signature"`
	Xmlns string `xml:"xmlns:ds,attr"`
	ID string `xml:"Id,attr,omitempty"`
	SignedInfo SignedInfo2
	SignatureValue SignatureValue2
	KeyInfo KeyInfo3
}
type EncryptionMethod2 struct {
	XMLName xml.Name `xml:"xenc:EncryptionMethod"`
	Algorithm string `xml:"Algorithm,attr,omitempty"`
}
type SignedInfo2 struct {
	XMLName xml.Name `xml:"ds:SignedInfo"`
	Xmlns string `xml:"xmlns:ds,attr,omitempty"`
	CanonicalizationMethod CanonicalizationMethod2
	SignatureMethod SignatureMethod2
	Reference []Reference3
}
type SignatureValue2 struct {
	XMLName xml.Name `xml:"ds:SignatureValue"`
	Contents []byte `xml:",innerxml"`
}
type CipherData2 struct {
	XMLName xml.Name `xml:"xenc:CipherData"`
	CipherValue CipherValue2
}
type ReferenceList2 struct {
	XMLName xml.Name `xml:"xenc:ReferenceList"`
	DataReference DataReference2
}
type CipherValue2 struct {
	XMLName xml.Name `xml:"xenc:CipherValue"`
	Contents []byte `xml:",innerxml"`
}
type DataReference2 struct {
	XMLName xml.Name `xml:"xenc:DataReference"`
	URI string `xml:"URI,attr,omitempty"`
}
type CanonicalizationMethod2 struct {
	XMLName xml.Name `xml:"ds:CanonicalizationMethod"`
	Algorithm string `xml:"Algorithm,attr,omitempty"`
}
type SignatureMethod2 struct {
	XMLName xml.Name `xml:"ds:SignatureMethod"`
	Algorithm string `xml:"Algorithm,attr,omitempty"`
}
type Reference3 struct {
	XMLName xml.Name `xml:"ds:Reference"`
	URI string `xml:"URI,attr,omitempty"`
	ValueType string `xml:"ValueType,attr,omitempty"`
	Transforms Transforms2
	DigestMethod DigestMethod2
	DigestValue DigestValue2
}
type Transforms2 struct {
	XMLName xml.Name `xml:"ds:Transforms"`
	Transform Transform2
}
type DigestMethod2 struct{
	XMLName xml.Name `xml:"ds:DigestMethod"`
	Algorithm string `xml:"Algorithm,attr,omitempty"`
}
type DigestValue2 struct{
	XMLName xml.Name `xml:"ds:DigestValue"`
	Contents []byte `xml:",innerxml"`
}
type Transform2 struct{
	XMLName xml.Name `xml:"ds:Transform"`
	Algorithm string `xml:"Algorithm,attr,omitempty"`
}
type KeyInfo2 struct{
	XMLName xml.Name `xml:"ds:KeyInfo"`
	Xmlns string `xml:"xmlns:ds,attr"`
	ID string `xml:"Id,attr,omitempty"`
	SecurityTokenReference SecurityTokenReference3
}
type KeyInfo3 struct{
	XMLName xml.Name `xml:"ds:KeyInfo"`
	ID string `xml:"Id,attr,omitempty"`
	SecurityTokenReference SecurityTokenReference2
}
type SecurityTokenReference3 struct {
	XMLName xml.Name `xml:"wsse:SecurityTokenReference"`
	Xmlns string `xml:"xmlns:wsse,attr,omitempty"`
	KeyIdentifier KeyIdentifier2
}
type KeyIdentifier2 struct {
	XMLName xml.Name `xml:"wsse:KeyIdentifier"`
	EncodingType string `xml:"EncodingType,attr,omitempty"`
	ValueType string `xml:"ValueType,attr,omitempty"`
	Contents []byte `xml:",innerxml"`
}
type SecurityTokenReference2 struct {
	XMLName xml.Name `xml:"wsse:SecurityTokenReference"`
	Xmlns string `xml:"xmlns:wsse,attr,omitempty"`
	WsuNs string `xml:"xmlns:wsu,attr,omitempty"`
	WsuId string `xml:"wsu:Id,attr,omitempty"`
	X509Data X509Data
}
type X509Data struct {
	XMLName xml.Name `xml:"ds:X509Data"`
	X509IssuerSerial X509IssuerSerial
}
type X509IssuerSerial struct {
	XMLName xml.Name `xml:"ds:X509IssuerSerial"`
	X509IssuerName X509IssuerName
	X509SerialNumber X509SerialNumber
}
type X509IssuerName struct {
	XMLName xml.Name `xml:"ds:X509IssuerName"`
	Contents   []byte `xml:",innerxml"`
}
type X509SerialNumber struct {
	XMLName xml.Name `xml:"ds:X509SerialNumber"`
	Contents   []byte `xml:",innerxml"`
}
type SignatureConfirmation struct {
	XMLName xml.Name `xml:"wsse11:SignatureConfirmation"`
	Xmlns string `xml:"xmlns:wsse11,attr"`
	WsuNs string `xml:"xmlns:wsu,attr"`
	Value string `xml:"Value,attr,omitempty"`
	WsuId string `xml:"wsu:Id,attr"`

}
type EncryptedData2 struct {
	XMLName xml.Name `xml:"xenc:EncryptedData"`
	Xmlns string `xml:"xmlns:xenc,attr,omitempty"`
	Id string `xml:"Id,attr,omitempty"`
	Type string `xml:"Type,attr,omitempty"`
	EncryptionMethod EncryptionMethod2
	KeyInfo KeyInfo4
	CipherData CipherData2
}
type KeyInfo4 struct {
	XMLName xml.Name `xml:"ds:KeyInfo"`
	Xmlns string `xml:"xmlns:ds,attr,omitempty"`
	SecurityTokenReference SecurityTokenReference4
}
type SecurityTokenReference4 struct {
	XMLName xml.Name `xml:"wsse:SecurityTokenReference"`
	Xmlns string `xml:"xmlns:wsse,attr,omitempty"`
	Reference Reference4
}
type Reference4 struct {
	XMLName xml.Name `xml:"wsse:Reference"`
	URI string `xml:"URI,attr,omitempty"`
}