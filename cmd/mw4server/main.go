package main

import (
	"fmt"
	"log"

	"github.com/Matovv/mw4server/internal/app/prefab"
	"github.com/Matovv/mw4server/internal/app/session"
	"github.com/Matovv/mw4server/internal/app/world"
	"github.com/Matovv/mw4server/internal/pkg/network"
	"github.com/Matovv/mw4server/internal/pkg/types"
)

const (
	PORT = 8880
)

func main() {
	log.Println("Magic Wars 4 - Server")
	sessionManager := session.NewSessionManager(10)
	log.Println("init session manager...")
	prefabManager := prefab.NewPrefabManager()
	prefabManager.InitTestData()
	world := world.NewWorld("test_1", prefabManager)
	world.SpawnUnit(1, types.FactionA, types.Vec2{X:0,Y:0})
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