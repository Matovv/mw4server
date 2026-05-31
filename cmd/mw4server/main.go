package main

import (
	"fmt"
	"log"

	"github.com/Matovv/mw4server/internal/app/session"
	"github.com/Matovv/mw4server/internal/pkg/network"
)

const (
	PORT = 8880
)

func main() {
	log.Println("Magic Wars 4 - Server")
	sessionManager := session.NewSessionManager(10)
	log.Println("init session manager...")
	testSession := session.NewSession(
		"session-1",
		"test",	
	)
	sessionManager.AddSession(testSession)
	go sessionManager.Run()
	server := network.NewServer(
		sessionManager,
	)
	log.Println("init server...")
	log.Println("port:", PORT)
	log.Fatal(server.Start(fmt.Sprintf(":%d", PORT)))
}