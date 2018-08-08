package main

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
)


//.CER — сертификат, или набор сертификатов, закодированных по стандарту CER.
//.DER — сертификат, закодированный по стандарту DER.
//.PEM — PEM-сертификат, закодированный по стандарту DER и использующий Base64 и помещенный между «----- BEGIN CERTIFICATE -----» и «----- END CERTIFICATE -----».
//.P7B, .P7C — PKCS #7 содержит несколько сертификатов или CRL.
//.P12 — PKCS #12 содержит блок, хранящий одновременно и закрытый ключ, и сертификат (в зашифрованном виде).
//.PFX — PFX, предшественник PKCS #12, также содержит блок закрытого ключа и сертификат.




func main() {
	http.HandleFunc("/", handler)
	// Start the HTTPS server in a goroutine
	//go http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil)
	// Start the HTTP server
	//log.Fatal(http.ListenAndServe(":8080", nil))
	log.Printf("start")
	http.ListenAndServe(":8088", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hi there!")
	//io.WriteString(w, "Hello world!")
	//	log.Printf("data recieved")
	body, _ := ioutil.ReadAll(r.Body)
	decryptedMessage,signatureText,_, smisErrorMessage, smisError := processXML(body)
	if smisError != nil {
		fmt.Println(smisErrorMessage)
	}
	fmt.Println(decryptedMessage)
	respText:=`<ns2:TestResponse xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns2="http://basis-edu.ru/monitoring/schemas/node" xmlns:xenc="http://www.w3.org/2001/04/xmlenc#"></ns2:TestResponse>`
	xmlResp, _ :=constructResponse(respText,string(signatureText))
	w.Header().Add("Content-Type", "text/xml; charset=utf-8")
	w.Header().Add("Accept","text/xml, text/html, image/gif, image/jpeg, *; q=.2, */*; q=.2")
	w.Header().Add("User-Agent","Java/1.7.0_80")
	w.Header().Add("Soapaction","\"http://basis-edu.ru/monitoring/schemas/node/TestResponse\"")
	ioutil.WriteFile("testResponseToSMIS.xml",[]byte(xmlResp),0644)
	fmt.Fprintf(w,xmlResp)
	//log.Printf(string(r.Body))
	defer r.Body.Close()
	//body, _ := ioutil.ReadAll(r.Body)
	//for k, v := range r.Header {
	//	fmt.Printf("Header field %q, Value %q\n", k, v)
	//}
	//_ = ioutil.WriteFile("response.xml", body, 0644)
	//log.Printf(r.Body)


}
