package main

import (
	"encoding/xml"
)

// AuthnRequest represents the SAML object of the same name, a request from a service provider
// to authenticate a user.
//
// See http://docs.oasis-open.org/security/saml/v2.0/saml-core-2.0-os.pdf
type SoapEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  SoapHeader
	Body    SoapBody
}
type SoapBody struct {
	XMLName       xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	EncryptedData EncryptedData
}
type SoapHeader struct {
	XMLName  xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Security wsSecurity
}
type wsSecurity struct {
	XMLName             xml.Name `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Security"`
	BinarySecurityToken BinarySecurityToken
	Signature			Signature
	EncryptedKey        EncryptedKey
}

type Signature struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Signature"`
	ID string `xml:"Id,attr,omitempty"`
	SignedInfo SignedInfo
	SignatureValue SignatureValue
	KeyInfo KeyInfo
}
type SignedInfo struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SignedInfo"`
	CanonicalizationMethod CanonicalizationMethod
	SignatureMethod SignatureMethod
	Reference Reference
}
type CanonicalizationMethod struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# CanonicalizationMethod"`
	Algorithm string `xml:"Algorithm,attr,omitempty"`
}
type SignatureMethod struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SignatureMethod"`
	Algorithm string `xml:"Algorithm,attr,omitempty"`
}
type Reference struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Reference"`
	URI string `xml:"URI,attr,omitempty"`
	ValueType string `xml:"ValueType,attr,omitempty"`
	Transforms Transforms
	DigestMethod DigestMethod
	DigestValue DigestValue
}
type Reference2 struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Reference"`
	URI string `xml:"URI,attr,omitempty"`
	ValueType string `xml:"ValueType,attr,omitempty"`
}
type Transforms struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Transforms"`
	Transform Transform
}
type DigestMethod struct{
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# DigestMethod"`
	Algorithm string `xml:"Algorithm,attr,omitempty"`
}
type DigestValue struct{
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# DigestValue"`
	Algorithm string `xml:"Algorithm,attr,omitempty"`
	Contents []byte `xml:",innerxml"`
}
type Transform struct{
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Transform"`
	Algorithm string `xml:"Algorithm,attr,omitempty"`
}
type SignatureValue struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SignatureValue"`
	Contents     []byte   `xml:",innerxml"`
}
type KeyInfo struct{
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# KeyInfo"`
	ID string `xml:"Id,attr,omitempty"`
	SecurityTokenReference SecurityTokenReference
}
type SecurityTokenReference struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd SecurityTokenReference"`
	KeyIdentifier KeyIdentifier
	Reference Reference2
}
type KeyIdentifier struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd KeyIdentifier"`
	EncodingType string `xml:"EncodingType,attr,omitempty"`
	ValueType string `xml:"ValueType,attr,omitempty"`
	Contents     []byte   `xml:",innerxml"`
}
type BinarySecurityToken struct {
	XMLName      xml.Name `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd BinarySecurityToken"`
	EncodingType string   `xml:"EncodingType,attr,omitempty"`
	ValueType string `xml:"ValueType,attr,omitempty"`
	Contents     []byte   `xml:",innerxml"`
}
type EncryptedKey struct {
	XMLName    xml.Name `xml:"http://www.w3.org/2001/04/xmlenc# EncryptedKey"`
	ID         string   `xml:"Id,attr,omitempty"`
	EncryptionMethod EncryptionMethod
	KeyInfo KeyInfo
	CipherData CipherData
	ReferenceList ReferenceList
}
type ReferenceList struct {
	XMLName    xml.Name `xml:"http://www.w3.org/2001/04/xmlenc# ReferenceList"`
	DataReference DataReference
}
type DataReference struct {
	XMLName    xml.Name `xml:"http://www.w3.org/2001/04/xmlenc# DataReference"`
	URI string `xml:"URI,attr,omitempty"`
}
type EncryptionMethod struct {
	XMLName    xml.Name `xml:"http://www.w3.org/2001/04/xmlenc# EncryptionMethod"`
	Algorithm string `xml:"Algorithm,attr,omitempty"`
}
type EncryptedData struct {
	XMLName    xml.Name `xml:"http://www.w3.org/2001/04/xmlenc# EncryptedData"`
	ID         string   `xml:"Id,attr,omitempty"`
	Type string `xml:"Type,attr,omitempty"`
	EncryptionMethod EncryptionMethod
	KeyInfo KeyInfo
	CipherData CipherData
}
type CipherData struct {
	XMLName     xml.Name `xml:"http://www.w3.org/2001/04/xmlenc# CipherData"`
	CipherValue CipherValue
}
type CipherValue struct {
	XMLName  xml.Name `xml:"http://www.w3.org/2001/04/xmlenc# CipherValue"`
	Contents []byte   `xml:",innerxml"`
}
type SoapBodyDecrypted struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ SOAP-ENV:Body"`
	Contents []byte  `xml:",innerxml"`
}