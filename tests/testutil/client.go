package testutil

import (
	"encoding/json"
	"testing"

	"github.com/Matovv/mw4server/internal/pkg/network"
	"github.com/gorilla/websocket"
)

type TestClient struct {
	Conn   *websocket.Conn
	UnitID uint64
}

func NewClient(
	t *testing.T,
	addr string,
) *TestClient {

	t.Helper()

	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		t.Fatalf("failed connect: %v", err)
	}

	return &TestClient{
		Conn: conn,
	}
}

func (c *TestClient) Close() {
	_ = c.Conn.Close()
}

func (c *TestClient) JoinSession(
	sessionID string,
) error {

	err := c.Conn.WriteJSON(
		map[string]any{
			"type": "join_session",
			"data": map[string]any{
				"session_id": sessionID,
			},
		},
	)

	if err != nil {
		return err
	}

	_, data, err := c.Conn.ReadMessage()
	if err != nil {
		return err
	}

	var resp network.JoinSessionResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return err
	}

	c.UnitID = resp.UnitID

	return nil
}

func (c *TestClient) Move(
	x int64,
	y int64,
) error {

	return c.Conn.WriteJSON(
		map[string]any{
			"type": "move",
			"data": map[string]any{
				"unit_id": c.UnitID,
				"x":       x,
				"y":       y,
			},
		},
	)
}