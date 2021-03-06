package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
	"crypto/sha1"
	"bytes"
	"crypto"
)



func processXML(byteXML []byte) (string,[]byte,[]byte,string,error) {
	soapEnvStructure := SoapEnvelope{} // init a structure from schema.go fully designed for the SOAP wsse input from SMIS object. not mutable for other SOAP inputs
	smisErrorMessage := ""  // init separately my error message and message from libs
	var smisError error
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error message: %s\n",smisErrorMessage)
			fmt.Printf("Error: %s\n",smisError)
		}
	}()
	// unmarshal xml - fills the fields of soapEnvStructure from xmlfile ([]byte required)
	if err := xml.Unmarshal(byteXML, &soapEnvStructure); err != nil {
		smisErrorMessage = "wrong XML format"
		smisError = err
		panic(smisErrorMessage)
	}
	// aes key is in the header inside encrypted key tag
	aesKeyEncrypted := soapEnvStructure.Header.Security.EncryptedKey.CipherData.CipherValue.Contents
	// check if did not find a key. Probably because the xml structure does not fit the structure from schema.go
	if len(aesKeyEncrypted)<1 {
		smisErrorMessage = "could not find aes key"
		smisError = fmt.Errorf("%s",smisErrorMessage)
		panic(smisErrorMessage)
	}
	// read private key from file, parse it (pem.decode) and make a x509 private key structure
	privateKeyData, err := ioutil.ReadFile("privateKey.pem")
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
	ciphertextb64 := []byte(strings.Replace(string(soapEnvStructure.Body.EncryptedData.CipherData.CipherValue.Contents), "\t", "", -1))
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
	signature := soapEnvStructure.Header.Security.Signature.SignatureValue.Contents
	signatureb64:=[]byte(strings.Replace(strings.Replace(string(signature), "\t", "", -1),"\r\n","",-1))
	signatureb64Decoded := make([]byte, base64.StdEncoding.DecodedLen(len(signatureb64)))
	n2,_ := base64.StdEncoding.Decode(signatureb64Decoded, signatureb64)
	digest := soapEnvStructure.Header.Security.Signature.SignedInfo.Reference.DigestValue.Contents
	idx:= strings.LastIndex(string(ciphertext),">")
	decryptedBody := changeBodyToDecrypted(string(ciphertext)[:idx+1],byteXML)
	decryptedBodyCanon,_ := ExcC14N(decryptedBody)
	h := sha1.New()
	h.Write(decryptedBodyCanon)
	digestFromBody := h.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(digestFromBody))
	fmt.Println(string(digest))
	if bytes.Compare(digest,digestFromBody)!=0 {
		//panic("Digest wrong!")
	}
	signedInfo,_ := getSignedInfo(byteXML)
	signedInfoCanon,_ := ExcC14N([]byte(signedInfo))
	//signedInfoCanon,_ = ioutil.ReadFile("signedInfoCanon.bin")
	h2 := sha1.New()
	h2.Write(signedInfoCanon)
	digestFromSignedInfo := h2.Sum(nil)

	publicKeyData, _ := ioutil.ReadFile("publicKeyObject.pem")
	blockPub, _ := pem.Decode([]byte(publicKeyData))
	rsaPub, _ := x509.ParsePKIXPublicKey(blockPub.Bytes)
	rsaPublicKey, _ := rsaPub.(*rsa.PublicKey)
	fmt.Println(base64.StdEncoding.EncodeToString(digestFromSignedInfo))
	err = rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA1, digestFromSignedInfo,signatureb64Decoded[:n2])
	//_ = ioutil.WriteFile("signedInfoCanon.bin",[]byte(signedInfoCanon),0644)
	return string(ciphertext)[:idx+1],signatureb64,digest, smisErrorMessage, smisError
}
