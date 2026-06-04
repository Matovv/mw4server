package testutil

import (
	"net"
	"testing"
	"time"

	"github.com/Matovv/mw4server/internal/app/prefab"
	"github.com/Matovv/mw4server/internal/app/session"
	"github.com/Matovv/mw4server/internal/app/world"
	"github.com/Matovv/mw4server/internal/pkg/network"
	"github.com/Matovv/mw4server/internal/pkg/types"
)

type TestService struct {
	SessionManager *session.SessionManager
	Server         *network.Server
	Addr           string
}

func StartTestService(
	t *testing.T,
) *TestService {

	t.Helper()

	prefabManager := prefab.NewPrefabManager()
	prefabManager.InitTestData()

	sessionManager := session.NewSessionManager(10)

	testSession := session.NewSession(
		"session-1",
		"test-map",
		world.NewWorld(
			"test-map",
			prefabManager,
		),
	)

	sessionManager.AddSession(testSession)

	server := network.NewServer(sessionManager)

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to allocate port: %v", err)
	}

	addr := listener.Addr().String()
	listener.Close()

	go func() {
		_ = server.Start(addr)
	}()

	deadline := time.Now().Add(2 * time.Second)

	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout(
			"tcp",
			addr,
			50*time.Millisecond,
		)
		if err == nil {
			_ = conn.Close()
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	return &TestService{
		SessionManager: sessionManager,
		Server:         server,
		Addr:           "ws://" + addr + "/ws",
	}
}

func (s *TestService) TickSession(
	t *testing.T,
	sessionID types.SessionID,
) *session.Session {

	t.Helper()

	session, err := s.SessionManager.GetSession(sessionID)
	if err != nil {
		t.Fatal(err)
	}

	session.SessionTick(s.SessionManager.TickRate)

	return session
}