	package main

import (
	"./MicroServs"
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