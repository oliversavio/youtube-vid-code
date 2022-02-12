package main

import (
	"log"

	_ "oliversavio/api/docs"
	"oliversavio/api/server"
)

// @title API Examples
// @version 1.0
// @description This is an auto-generated API Docs.
// @contact.name API Support
// @contact.email your@mail.com
// @BasePath /api
func main() {

	svr := &server.Server{
		App: server.InitFiberApplication(),
	}

	waitforShutdownInterrupt := svr.Start()
	<-waitforShutdownInterrupt

	log.Println("Shutting Down Server..")
	svr.ShutdownGracefully()
}
