package network

import (
	"encoding/json"
	"sync/atomic"

	"github.com/Matovv/mw4server/internal/pkg/errors"
	"github.com/Matovv/mw4server/internal/pkg/types"
	"github.com/gorilla/websocket"
)

type Client struct {
    ID uint64
    Conn *websocket.Conn
    SessionID types.SessionID
    SendChan chan []byte
    Closed atomic.Bool
}

func (c *Client) Close() {
	if !c.Closed.CompareAndSwap(
		false,
		true,
	) {
		return
	}
	_ = c.Conn.Close()
}

func (c *Client) GetId() uint64 {
	return c.ID
}

func (c *Client) SendPacket(
	packetType string,
	payload any,
) error {
	if c.Closed.Load() {
		return errors.ErrDisconnected
	}
	data, err := NewPacket(
		packetType,
		payload,
	)
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
	for {
		msg, ok := <-c.SendChan
		if !ok {
			return
		}
		err := c.Conn.WriteMessage(
			websocket.TextMessage,
			msg,
		)
		if err != nil {
			c.Close()
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