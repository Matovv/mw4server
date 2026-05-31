package network

import (
	"encoding/json"
	"log"

	"github.com/Matovv/mw4server/internal/app/command"
	"github.com/Matovv/mw4server/internal/pkg/types"
)

func (s *Server) handleJoinSession(
	client *Client,
	raw json.RawMessage,
) error {
	var req JoinSessionRequest	
	if err := json.Unmarshal(raw, &req); err != nil {
		return err
	}
	//log.Println("data:", req)
	session, err := s.SessionManager.GetSession(
		types.SessionID(req.SessionID),
	)
	if err != nil {
		return err
	}		
	client.SessionID = types.SessionID(req.SessionID)	
	slot, err := session.AssignPlayer(
		types.PlayerID(client.ID),
		client,
	)
	if err != nil {
		return err
	}
	response := JoinSessionResponse{
		PlayerID:  client.ID,
		SlotIndex: slot.Index,
		Tick:      session.GetTick(),
	}	
	client.Send(response)
	return nil
}

func (s *Server) handleMove(
    client *Client,
    raw json.RawMessage,
) error {
    var req MoveRequest
    if err := json.Unmarshal(raw, &req); err != nil {
        return err
    }
	log.Println("data:", req)
    session, err := s.SessionManager.GetSession(
		types.SessionID(client.SessionID),
	)
	if err != nil {
		return err
	}	
    session.CommandPush(
        &command.MoveCommand{
            PlayerID: types.PlayerID(client.ID),
            UnitID: types.UnitID(req.UnitID),
            Position: types.Vec2{
                X: req.X,
                Y: req.Y,
            },
        },
    )
	return nil
}