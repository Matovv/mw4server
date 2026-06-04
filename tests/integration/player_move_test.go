package integration

import (
	"testing"
	"time"

	"github.com/Matovv/mw4server/internal/pkg/types"
	"github.com/Matovv/mw4server/tests/testutil"
)

func TestPlayerMove(
	t *testing.T,
) {
	svc := testutil.StartTestService(t)
	client := testutil.NewClient(
		t,
		svc.Addr,
	)
	defer client.Close()
	err := client.JoinSession("session-1")
	if err != nil {
		t.Fatal(err)
	}
	err = client.Move(-1000, 0)
	if err != nil {
		t.Fatal(err)
	}
	session, err := svc.SessionManager.GetSession("session-1")
	if err != nil {
		t.Fatal(err)
	}
	ok := testutil.WaitUntil(
		func() bool {
			return session.CommandQueue.Len() > 0
		},
		1 * time.Second,
	)
	if !ok {
		t.Fatal("move command never reached server")
	}
	session = svc.TickSession(t, "session-1")
	unit := session.World.Units[
		types.UnitID(client.UnitID),
	]
	if unit.Transform.Position.X == 0 {
		t.Fatal("unit did not move")
	}
	if unit.Transform.Rotation == 0 {
		t.Fatal("unit did not rotate")
	}
}