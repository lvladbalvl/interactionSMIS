package smisInteract

import (
	"encoding/xml"
	"math/rand"
	"fmt"
	"crypto/sha1"
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
	"encoding/pem"
	"crypto/rsa"
	"crypto/x509"
	crytporand "crypto/rand"
	"crypto"
	"bytes"
	"io"
)

const pathToPrivateKey string = "./privateKey.pem"

func ConstructResponse(respText string, signature string,rsaPublicKey *rsa.PublicKey) (string,error) {

	var soapBodyDigest []byte
	var signatureResp []byte
	var aesKeyEncryptedB64 string
	var soapBodyEncrypted []byte
	// let's construct envelope of the response
	soapResp := SoapEnvelope2{}
	soapResp.Xmlns = "http://schemas.xmlsoap.org/soap/envelope/"
	soapResp.Header.Security.Xmlns = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	soapResp.Header.Security.MustUnderstand = "1"
	soapResp.Header.Security.EncryptedKey.Xmlns = "http://www.w3.org/2001/04/xmlenc#"
	soapResp.Header.Security.EncryptedKey.Id = "EncKeyId-C5619879ACACCFA2FB152827572172050155"
	soapResp.Header.Security.EncryptedKey.EncryptionMethod.Algorithm = "http://www.w3.org/2001/04/xmlenc#rsa-1_5"
	soapResp.Header.Security.EncryptedKey.KeyInfo.Xmlns = "http://www.w3.org/2000/09/xmldsig#"
	soapResp.Header.Security.EncryptedKey.CipherData.CipherValue.Contents = []byte(aesKeyEncryptedB64)
	soapResp.Header.Security.EncryptedKey.ReferenceList.DataReference.URI = "#EncDataId-"+fmt.Sprintf("%05d",rand.Intn(9999))
	soapResp.Header.Security.EncryptedKey.KeyInfo.SecurityTokenReference.Xmlns = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	soapResp.Header.Security.EncryptedKey.KeyInfo.SecurityTokenReference.KeyIdentifier.EncodingType = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary"
	soapResp.Header.Security.EncryptedKey.KeyInfo.SecurityTokenReference.KeyIdentifier.ValueType = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-x509-token-profile-1.0#X509v3"
	soapResp.Header.Security.EncryptedKey.KeyInfo.SecurityTokenReference.KeyIdentifier.Contents = []byte("MIIDWTCCAkGgAwIBAgIRAIf39GOHSk0bvUWn/N0sGskwDQYJKoZIhvcNAQENBQAwIjELMAkGA1UEBhMCUlUxEzARBgNVBAMMCtCh0YLQtdC90LQwHhcNMTgwNDIwMDUyMTMzWhcNMTkxMjMxMjEwMDAwWjAiMQswCQYDVQQGEwJSVTETMBEGA1UEAwwK0KHRgtC10L3QtDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAJkrlaphIDd6Q63WOvwJZ4nckllRu8w0yg0IEl8fi6PS7nv1xwND4rQiVPHb08v+aVNpxC6+Lthx97D6qz8qaJc0zS/EDV8NY+VTrbXAZbxoJT4oLtE2z8uszaUEDtQCGlz79qcCfsGeSulOyfHXlJWMLy/zXPFHwvfcoL2iM+NOSo12Icw9etehLNCm5pOZ4INwQj0PnJ6rzC4epjf+j8U8tF+oJhb0DrQyihqgojMIJNe5wX2iGerA8NOcgiUWq6cSgJH0t/lePO3mcQgUauje3EsnCuLIYWiMH20WY08Z3xYNn33WKeWfK2mSFlrYf6gKIz+aWG+oEmDxgEfHaH8CAwEAAaOBiTCBhjAMBgNVHRMBAf8EAjAAMCAGA1UdDgEB/wQWBBThGTpSygr8MMwAhOFOnwH6NfL0xzAiBgNVHSMBAf8EGDAWgBThGTpSygr8MMwAhOFOnwH6NfL0xzAOBgNVHQ8BAf8EBAMCBLAwIAYDVR0lAQH/BBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMBMA0GCSqGSIb3DQEBDQUAA4IBAQBd/IaZleMlR4QbWX7e0iuJvbyJ6Gid4wVOxo8ckXwncbnpR/02QrnY7w3WTiqZb8SNYz9jjODHXlozxwTiSTQBbxqxz1dDM3K2WOIL8YeOO0xLIddJfnkOkrcUXDim2eCfMe9jBxoG27AlIfWkzCYC3yGkqLxjebohEjRww5/5s3dk0N0eJuRBEgfpRabCu3X4QBNEfaYQZHX43foofjWGVRw9keHBYlNTTYjuSG0G88ITDu++dlQdwDvWrDpavdABp780aM375y6q4wvDUCCIgXtaXwbBbyAu73Hi0pgKJ+Lt624xERwTYxANBwHwPhsaLp2m5qQ8XHcPmf90lDtc")
	soapResp.Header.Security.Signature.Xmlns = "http://www.w3.org/2000/09/xmldsig#"
	soapResp.Header.Security.Signature.ID = "Signature-"+fmt.Sprintf("%05d",rand.Intn(9999))
	soapResp.Header.Security.Signature.SignedInfo.CanonicalizationMethod.Algorithm = "http://www.w3.org/2001/10/xml-exc-c14n#"
	soapResp.Header.Security.Signature.SignedInfo.SignatureMethod.Algorithm = "http://www.w3.org/2000/09/xmldsig#rsa-sha1"
	soapResp.Header.Security.Signature.KeyInfo.ID = "KeyId-C5619879ACACCFA2FB152827572170250152"
	soapResp.Header.Security.Signature.KeyInfo.SecurityTokenReference.Xmlns = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	soapResp.Header.Security.Signature.KeyInfo.SecurityTokenReference.WsuNs = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	soapResp.Header.Security.Signature.KeyInfo.SecurityTokenReference.WsuId = "STRId-C5619879ACACCFA2FB152827572170250153"
	soapResp.Header.Security.Signature.KeyInfo.SecurityTokenReference.X509Data.X509IssuerSerial.X509IssuerName.Contents = []byte("CN=ЕДДС,C=RU")
	soapResp.Header.Security.Signature.KeyInfo.SecurityTokenReference.X509Data.X509IssuerSerial.X509SerialNumber.Contents = []byte("297708304065041750947242433498698901951")
	// then construct signature confirmation separately - only one tag with the signature of the received message.
	// To ensure that this is the response to the message recieved
	// also we will put the digest from it inside the signed info
	sigConf,digestFromSigConf := processSigConf(signature)
	// now encrypt message body
	soapBodyEncrypted, soapBodyDigest, aesKeyEncryptedB64,soapBody,EncrDataByte := encryptMsg(soapResp,rsaPublicKey,respText)
	soapBody.Contents = EncrDataByte
	soapResp.Body = soapBody
	// fill in all rematinin fields of the signed info and then sign it
	soapResp.Header.Security.SignatureConfirmation = sigConf
	soapResp.Header.Security.Signature.SignatureValue.Contents = []byte(base64.StdEncoding.EncodeToString(signatureResp))
	soapResp.Header.Security.Signature.SignedInfo.Reference = make([]Reference3,2)
	soapResp.Header.Security.Signature.SignedInfo.Reference[0].URI = "#"+soapBody.WsuId
	soapResp.Header.Security.Signature.SignedInfo.Reference[0].Transforms.Transform.Algorithm = "http://www.w3.org/2001/10/xml-exc-c14n#"
	soapResp.Header.Security.Signature.SignedInfo.Reference[0].DigestMethod.Algorithm = "http://www.w3.org/2000/09/xmldsig#sha1"
	soapResp.Header.Security.Signature.SignedInfo.Reference[0].DigestValue.Contents = []byte(base64.StdEncoding.EncodeToString(soapBodyDigest))
	soapResp.Header.Security.Signature.SignedInfo.Reference[1].URI = "#"+sigConf.WsuId
	soapResp.Header.Security.Signature.SignedInfo.Reference[1].Transforms.Transform.Algorithm = "http://www.w3.org/2001/10/xml-exc-c14n#"
	soapResp.Header.Security.Signature.SignedInfo.Reference[1].DigestMethod.Algorithm = "http://www.w3.org/2000/09/xmldsig#sha1"
	soapResp.Header.Security.Signature.SignedInfo.Reference[1].DigestValue.Contents = []byte(base64.StdEncoding.EncodeToString(digestFromSigConf))
	// generating signature via signed info digest. signed info will be taken from soapResp
	signatureResp = genSignature(soapResp,soapBodyEncrypted)

	soapRespString,_ := xml.MarshalIndent(soapResp,"","")
	return string(soapRespString),nil
}

func Pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}
func encrypt(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	msg := Pad([]byte(text))
	ciphertext := make([]byte, aes.BlockSize+len(msg))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(crytporand.Reader, iv); err != nil {
		return "", err
	}

	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(ciphertext[aes.BlockSize:], []byte(msg))
	return string(ciphertext), nil
}
func processSigConf(signature string) (sigConf SignatureConfirmation, digestFromSigConf []byte) {
	sigConf.Value = signature
	sigConf.WsuNs = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	sigConf.Xmlns = "http://docs.oasis-open.org/wss/oasis-wss-wssecurity-secext-1.1.xsd"
	sigConf.WsuId = "SigConf-"+fmt.Sprintf("%05d",rand.Intn(9999))
	sigConfString,_ := xml.MarshalIndent(sigConf,"","")
	sigConfStringC14N := ExcC14N(sigConfString)
	h := sha1.New()
	h.Write(sigConfStringC14N)
	digestFromSigConf = h.Sum(nil)
	return
}
func encryptMsg(soapResp SoapEnvelope2,rsaPublicKey *rsa.PublicKey,respText string) (soapBodyEncrypted []byte ,soapBodyDigest []byte, aesKeyEncryptedB64 string, soapBody SoapBody2,EncrDataByte []byte) {
	aesKey := make([]byte, 16)
	rand.Read(aesKey)
	aesKeyEncrypted,_ := rsa.EncryptPKCS1v15(crytporand.Reader,rsaPublicKey,aesKey)
	aesKeyEncryptedB64 = base64.StdEncoding.EncodeToString(aesKeyEncrypted)
	soapBody.WsuNs = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	soapBody.WsuId = "id-"+fmt.Sprintf("%05d",rand.Intn(9999))
	soapBody.Contents = []byte(respText)
	soapBodyByte,_ := xml.MarshalIndent(soapBody,"","")
	soapBodyByteCanoned := ExcC14N(soapBodyByte)
	h2 := sha1.New()
	h2.Write(soapBodyByteCanoned)
	soapBodyDigest = h2.Sum(nil)
	soapBodyEncryptedByte,_ := encrypt(aesKey,respText)
	soapBodyEncrypted = []byte(soapBodyEncryptedByte)
	EncrData := EncryptedData2{}
	EncrData.Xmlns = "http://www.w3.org/2001/04/xmlenc#"
	EncrData.Id = soapResp.Header.Security.EncryptedKey.ReferenceList.DataReference.URI[1:]
	EncrData.Type = "http://www.w3.org/2001/04/xmlenc#Content"
	EncrData.EncryptionMethod.Algorithm = "http://www.w3.org/2001/04/xmlenc#aes128-cbc"
	EncrData.KeyInfo.Xmlns = "http://www.w3.org/2000/09/xmldsig#"
	EncrData.KeyInfo.SecurityTokenReference.Xmlns = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	//maybe generate Id in separate variable
	EncrData.KeyInfo.SecurityTokenReference.Reference.URI = soapResp.Header.Security.EncryptedKey.Id
	EncrData.CipherData.CipherValue.Contents = []byte(base64.StdEncoding.EncodeToString(soapBodyEncrypted))
	EncrDataByte,_ = xml.Marshal(EncrData)
	return
}
func genSignature(soapResp SoapEnvelope2,soapBodyEncrypted []byte) (signatureResp []byte) {
	privateKeyData, _ := ioutil.ReadFile(pathToPrivateKey)
	privateKeyBlock, _ := pem.Decode(privateKeyData)
	var pri *rsa.PrivateKey
	pri, _ = x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	signedInfo,_ := xml.MarshalIndent(soapResp.Header.Security.Signature.SignedInfo,"","")
	signedInfoCanoned := ExcC14N(signedInfo)
	h3 := sha1.New()
	h3.Write(signedInfoCanoned)
	digestFromSignedInfo := h3.Sum(nil)
	signatureResp,_ = rsa.SignPKCS1v15(crytporand.Reader, pri, crypto.SHA1, digestFromSignedInfo)
	return
}