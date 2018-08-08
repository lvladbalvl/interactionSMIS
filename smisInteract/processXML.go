package smisInteract

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"strings"
	"crypto/sha1"
	"crypto"
	"regexp"
)



func ProcessXML(byteXML []byte) (string,[]byte,[]byte,string,error) {
	//soapEnvStructure := SoapEnvelope{} // init a structure from schema.go fully designed for the SOAP wsse input from SMIS object. not mutable for other SOAP inputs
	var smisErrorMessage string  // init separately my error message and message from libs
	var smisError error
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error message: %s\n",smisErrorMessage)
			fmt.Printf("Error: %s\n",smisError)
		}
	}()

	// unmarshal xml - fills the fields of soapEnvStructure from xmlfile ([]byte required)
	//if err := xml.Unmarshal(byteXML, &soapEnvStructure); err != nil {
	//	smisErrorMessage = "wrong XML format"
	//	smisError = err
	//	panic(smisError)
	//}
	// check digest and signature

	digest,signature,certBytes,aesKeyEncrypted,cipheredText := getNecessaryFields(byteXML)
	//digest := soapEnvStructure.Header.Security.Signature.SignedInfo.Reference.DigestValue.Contents
	//signature := soapEnvStructure.Header.Security.Signature.SignatureValue.Contents
	//get certificate from message and get public key from it
	//certBytes :=soapEnvStructure.Header.Security.BinarySecurityToken.Contents
	//aesKeyEncrypted := soapEnvStructure.Header.Security.EncryptedKey.CipherData.CipherValue.Contents
	//cipheredText := soapEnvStructure.Body.EncryptedData.CipherData.CipherValue.Contents
	certBlock := make([]byte, base64.StdEncoding.DecodedLen(len(certBytes)))
	n, err := base64.StdEncoding.Decode(certBlock, certBytes)
	cert,err := x509.ParseCertificate(certBlock[:n])
	if err!=nil {
		smisErrorMessage = "Wrong certificate in message!"
		smisError = fmt.Errorf("%s",smisErrorMessage)
		panic(smisError)
	}
	//fmt.Println(cert.SignatureAlgorithm.String())
	//err =cert.CheckSignatureFrom(cert)
	//fmt.Println(err)
	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)
	err =cert.CheckSignature(cert.SignatureAlgorithm,cert.RawTBSCertificate,cert.Signature)
	fmt.Println(err)
	signatureb64,smisErrorMessage,smisError:=checkDigestAndSignature(signature,digest,byteXML,rsaPublicKey)
	if smisError!= nil {
		panic(smisError)
	}

	// aes key is in the header inside encrypted key tag

	// check if did not find a key.

	if len(aesKeyEncrypted)<1 {
		smisErrorMessage = "could not find aes key"
		smisError = fmt.Errorf("%s",smisErrorMessage)
		panic(smisErrorMessage)
	}


	messageDecrypted,smisErrorMessage,smisError := decryptSoapBody(aesKeyEncrypted,cipheredText)
	return messageDecrypted,signatureb64,certBytes, smisErrorMessage, smisError
}


	func checkDigestAndSignature(signature []byte,digest []byte, byteXML []byte,rsaPublicKey *rsa.PublicKey) (signatureb64 []byte, smisErrorMessage string, smisError error) {
		// get signature and b64 decode
		signatureb64=[]byte(strings.Replace(strings.Replace(string(signature), "\t", "", -1),"\r\n","",-1))
		signatureb64Decoded := make([]byte, base64.StdEncoding.DecodedLen(len(signatureb64)))
		n2,_ := base64.StdEncoding.Decode(signatureb64Decoded, signatureb64)
		// get digest and b64 decode. SUITABLE FOR ONLY ONE DIGEST PER MESSAGE. IF THE MESSAGE WOULD BE DIFFERENT - NEED TO REWRITE CODE TO USE REFERENCES

		// get soap body. the digest is calculated from the encrypred soap-body.
		cipheredBody,smisError := getSoapBody(byteXML)
		if smisError!=nil {
			smisErrorMessage = "Wrong XML format"
			panic(smisError)
		}
		// xml-exc-c14n first and then compute hash and compare
		cipheredBodyCanon := ExcC14N(cipheredBody)
		h := sha1.New()
		h.Write(cipheredBodyCanon)
		digestFromBody := h.Sum(nil)
		if strings.Compare(string(digest),base64.StdEncoding.EncodeToString(digestFromBody))!=0 {
			smisErrorMessage = "Digest wrong!"
			smisError = fmt.Errorf("%s",smisErrorMessage)
			panic(smisError)
		}
		// now get signed info. the signature is checked first via the digest of canonicalized signed info
		signedInfo,smisError := getSignedInfo(byteXML)
		if smisError!=nil {
			smisErrorMessage = "Wrong XML format"
			panic(smisError)
		}
		signedInfoCanon := ExcC14N([]byte(signedInfo))
		h2 := sha1.New()
		h2.Write(signedInfoCanon)
		digestFromSignedInfo := h2.Sum(nil)

		// verify signature
		err := rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA1, digestFromSignedInfo,signatureb64Decoded[:n2])
		if err!=nil {
			smisErrorMessage = "Signature wrong!"
			smisError = err
			panic(smisError)
		}
		return
	}

	func decryptSoapBody(aesKeyEncrypted []byte,cipheredText []byte) (message string, smisErrorMessage string, smisError error) {

		// read private key from file, parse it (pem.decode) and make a x509 private key structure
		privateKeyData, err := ioutil.ReadFile(pathToPrivateKey)
		privateKeyBlock, _ := pem.Decode(privateKeyData)
		var pri *rsa.PrivateKey
		pri, err = x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
		if err != nil {
			smisErrorMessage ="something wrong with private key"
			smisError = err
			panic(smisErrorMessage)
		}
		//base64 decoding for []byte is made to the buffer => init the same size as needed
		aesKeyEncryptedDecoded := make([]byte, base64.StdEncoding.DecodedLen(len(aesKeyEncrypted)))
		n, err := base64.StdEncoding.Decode(aesKeyEncryptedDecoded, aesKeyEncrypted)
		// don't know why but the base64 decode mistakes by n bytes - but outputs this value which is then used to limit the slices
		if err != nil {
			smisErrorMessage ="something wrong with base64 format of aes key in XML"
			smisError = err
			panic(smisErrorMessage)
		}
		// decrypt aes key by rsa. the key is limited by n - see above
		aesKeyDecrypted, err := rsa.DecryptPKCS1v15(rand.Reader, pri, aesKeyEncryptedDecoded[:n])
		if err != nil {
			smisErrorMessage ="can not decrypt aes key"
			smisError = err
			panic(smisErrorMessage)
		}
		// xml unmarshals encrypted body with tab inside - so let's make it string, replace tabs and make []byte again
		ciphertextb64 := []byte(strings.Replace(string(cipheredText), "\t", "", -1))
		// base64 decoding after the tabs were deleted. the comments concerning buffer and n mistake value are applicable here as well (see above)
		ciphertext := make([]byte, base64.StdEncoding.DecodedLen(len(ciphertextb64)))
		n, err = base64.StdEncoding.Decode(ciphertext, ciphertextb64)
		if err != nil {
			smisErrorMessage ="something wrong with base64 format of encrypted body in XML"
			smisError = err
			panic(smisErrorMessage)
		}
		// aes128-cbc needs the initialization vector (iv) of 16 bytes. usually is in the beginning of the data. let's split the data into the iv and the data itself
		iv := ciphertext[:aes.BlockSize]
		ciphertext = ciphertext[aes.BlockSize:n]
		// construct aes cipher from key decrypted previously
		block, err := aes.NewCipher(aesKeyDecrypted)
		if err != nil {
			smisErrorMessage ="aes key not applicable"
			smisError = err
			panic(smisErrorMessage)
		}
		// the length os the message must be divisible by block size (16 byte) because the decryption is made by blocks.
		// if the data is changed but the number of blocks remains the same - no error will happen.
		// it would just result in rubbish
		if len(ciphertext)%aes.BlockSize != 0 {
			smisErrorMessage = "ciphertext is not a multiple of the block size"
			panic(smisErrorMessage)
		}
		//let's decrypt. if the block size is correct - no error will be expected
		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptBlocks(ciphertext, ciphertext)
		idx:= strings.LastIndex(string(ciphertext),">")
		message = string(ciphertext[:idx+1])
		return
	}
	func getNecessaryFields(byteXML []byte) (digest []byte, signature []byte, certificate []byte, aesKey []byte, cipherText []byte) {
		r := regexp.MustCompile(`<ds:DigestValue((.|\r\n)*?)</ds:DigestValue>`)
		res := r.FindAll(byteXML,-1)
		digest = getTokenContent(res[0])
		r = regexp.MustCompile(`<ds:SignatureValue((.|\r\n)*?)</ds:SignatureValue>`)
		res = r.FindAll(byteXML,-1)
		signature = getTokenContent(res[0])

		r = regexp.MustCompile(`<wsse:BinarySecurityToken((.|\r\n)*?)</wsse:BinarySecurityToken>`)
		res = r.FindAll(byteXML,-1)
		certificate = getTokenContent(res[0])
		r = regexp.MustCompile(`<xenc:CipherValue((.|\r\n)*?)</xenc:CipherValue>`)
		res = r.FindAll(byteXML,-1)
		aesKey = getTokenContent(res[0])
		cipherText = getTokenContent(res[1])
		return
	}
	func getTokenContent(token []byte) []byte {
		return []byte(strings.Replace(strings.Split(strings.Split(string(token),">")[1],"<")[0],"\r\n","",-1))
	}