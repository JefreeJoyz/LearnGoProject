package main

import (
	"learn-go-project/services"
	"net/http"
)

func main() {

	//mainFunc("Мир")
	//showMessage()
	//enterTheClub(19)
	//check()
	//checkVariable()
	//printMessages(messa)
	//workWithSlices(messaSlice)
	//fForMatrix()
	//PanicRecoverDefer()
	//testMap()
	//testStruct()
	//testInterface()
	//services.GetMyIp()
	//testUnitTest()
	fileServer := http.FileServer(http.Dir("./services"))
	http.Handle("/", fileServer)
	http.HandleFunc("/handler1", services.GetMyIpEndpoint)
	http.ListenAndServe(":8080", nil)
	//workWithEnv()
}
