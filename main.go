package main

import (
	_ "github.com/lib/pq"
	"fmt"
	"database/sql"
)
const (
	DB_USER     = "postgres"
	DB_PASSWORD = "admin"
	DB_NAME     = "test"
)
func main() {

	//xmlFile, _ := os.Open("response_marsh.xml")
	//defer xmlFile.Close()
	//byteXML, _ := ioutil.ReadAll(xmlFile)
	//t0:=time.Now()
	//encryptedText,signature,publicKey,_, _:= smisInteract.ProcessXML(byteXML)
	//t1:=time.Now()
	//fmt.Printf("Time passed:%s\r\n",t1.Sub(t0))
	//fmt.Println(encryptedText)
	//respText:=`<ns2:TestResponse xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns2="http://basis-edu.ru/monitoring/schemas/node" xmlns:xenc="http://www.w3.org/2001/04/xmlenc#"></ns2:TestResponse>`
	//xmlResp, _ :=smisInteract.ConstructResponse(respText,string(signature),publicKey)
	//t2:=time.Now()
	//fmt.Printf("Time passed:%s\r\n",t2.Sub(t0))
	//ioutil.WriteFile("soapEnvelopeMarshalTesting.xml",[]byte(xmlResp),0644)
	//nc,_ := nats.Connect(nats.DefaultURL)
	//nc.QueueSubscribe("foo","queue1", func(m *nats.Msg) {
	//	fmt.Printf("Received a message: %s\n", string(m.Data))
	//})
	//runtime.Goexit()

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, _ := sql.Open("postgres", dbinfo)
	fmt.Println(db)
	_, err := db.Exec(`INSERT INTO public.test4(
	name, surname, email)
	VALUES ( 'sdf', 'fgdsgsdf', 'dgggg');`);
	fmt.Println(err)
}