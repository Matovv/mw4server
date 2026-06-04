package integration

import (
	"testing"

	"github.com/Matovv/mw4server/tests/testutil"
)

func TestJoinSession(
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
}