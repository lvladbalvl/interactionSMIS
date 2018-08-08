package smisInteract

import (
	"strings"
	"fmt"
)

func ChooseAnswer(text []byte) string {
	// need to make error handling
	firstCar := strings.SplitN(string(text)," ",2)
	secCar:=strings.Split(firstCar[0],":")
	if len(secCar)!=2 {
		fmt.Println("Error!")
	}
	request:=secCar[1]
	response:=""
	switch request {
	case "DispatchMessageRequest":
		response=`<ns2:DispatchMessageResponse xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns2="http://basis-edu.ru/monitoring/schemas/node" xmlns:xenc="http://www.w3.org/2001/04/xmlenc#"></ns2:DispatchMessageResponse>`
	case "TestRequest":
		response=`<ns2:TestResponse xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns2="http://basis-edu.ru/monitoring/schemas/node" xmlns:xenc="http://www.w3.org/2001/04/xmlenc#"></ns2:TestResponse>`
	case "DispatchControlPointRequest":
		response=`<ns2:DispatchControlPointResponse xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns2="http://basis-edu.ru/monitoring/schemas/node" xmlns:xenc="http://www.w3.org/2001/04/xmlenc#"></ns2:DispatchControlPointResponse>`
	case "DispatchMaintenanceRequest":
		response=`<ns2:DispatchMaintenanceResponse xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns2="http://basis-edu.ru/monitoring/schemas/node" xmlns:xenc="http://www.w3.org/2001/04/xmlenc#"></ns2:DispatchMaintenanceResponse>`
	default:
		response="Unknown"
		}
		return response
}