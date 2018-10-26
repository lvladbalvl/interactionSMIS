	package main

import (
	"./microServs"
	"runtime"
)
func main() {
	go microServs.MsgHandler()
	go microServs.DbWorker()
	go microServs.MsgConstructor()
	//go microServs.MsgSender()
	go microServs.MsgReceiver()
	runtime.Goexit()
	}
