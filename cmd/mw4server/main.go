package main

import (
	"fmt"
	"log"

	"github.com/Matovv/mw4server/internal/app/prefab"
	"github.com/Matovv/mw4server/internal/app/session"
	"github.com/Matovv/mw4server/internal/app/world"
	"github.com/Matovv/mw4server/internal/pkg/network"
)

const (
	PORT = 8880
)

func main() {
	log.Println("Magic Wars 4 - Server")
	sessionManager := session.NewSessionManager(10)
	log.Println("init session manager...")
	prefabManager := prefab.NewPrefabManager()
	world := world.NewWorld("test_1", prefabManager)
	testSession := session.NewSession(
		"session-1",
		"test",	
		world,
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