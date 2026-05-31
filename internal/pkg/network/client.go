package network

import (
	"encoding/json"

	"github.com/Matovv/mw4server/internal/pkg/errors"
	"github.com/Matovv/mw4server/internal/pkg/types"
	"github.com/gorilla/websocket"
)

type Client struct {
    ID uint64

    Conn *websocket.Conn

    SessionID types.SessionID

    SendChan chan []byte
}
func (c *Client) GetId() uint64 {
	return c.ID
}

func (c *Client) Send(v any) error {

	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	select {
	case c.SendChan <- data:
		return nil

	default:
		return errors.ErrSendQueueFull
	}
}

func (c *Client) writeLoop() {

    defer c.Conn.Close()

    for msg := range c.SendChan {

        err := c.Conn.WriteMessage(
            websocket.TextMessage,
            msg,
        )

        if err != nil {
            return
        }
    }
}

func (c *Client) readLoop(
    server *Server,
) {

    defer c.Conn.Close()

    for {

        _, data, err := c.Conn.ReadMessage()

        if err != nil {
            return
        }

        var packet Packet

        if err := json.Unmarshal(
            data,
            &packet,
        ); err != nil {
            continue
        }

        server.HandlePacket(
            c,
            packet,
        )
    }
}