package session

import (
	"log"
	"time"

	"github.com/Matovv/mw4server/internal/app/command"
	"github.com/Matovv/mw4server/internal/app/protocol"
	"github.com/Matovv/mw4server/internal/app/session/models"
	"github.com/Matovv/mw4server/internal/app/world"
	"github.com/Matovv/mw4server/internal/pkg/errors"
	"github.com/Matovv/mw4server/internal/pkg/types"
)

type Session struct {
	ID types.SessionID
	MapName string
	Slots [8]*models.PlayerSlot
	Clients map[uint64]ClientSender
	World *world.World
	CommandQueue *CommandQueue
	Tick uint64
	StartedAt time.Time
}

func NewSession(
	id types.SessionID,
	mapName string,
) *Session {
	s := &Session{
		ID: id,
		MapName: mapName,
		Clients: make(map[uint64]ClientSender),
		CommandQueue: &CommandQueue{},
		World: world.NewWorld(mapName),
	}
	for i := range s.Slots {
		s.Slots[i] = &models.PlayerSlot{
			Index: uint8(i),
		}
	}
	return s
}

func (s *Session) GetTick() uint64 {
	return s.Tick
}

func (s *Session) AssignPlayer(
    playerID types.PlayerID,
	client ClientSender,
) (*models.PlayerSlot, error) {
	log.Println("assigning player...")
    for i, slot := range s.Slots {
        if slot == nil {
            continue
        }
        if !slot.Connected {
            slot.PlayerID = playerID
            slot.Connected = true
			s.Clients[client.GetId()] = client
			log.Println("player",playerID,"assigned to session", s.ID, "slot", i)
            return slot, nil
        }
    }
    return nil, errors.ErrSessionFull
}

func (s *Session) RemovePlayer(
	playerID types.PlayerID,
	force bool,
) {
	delete(s.Clients, uint64(playerID))
	for i, slot := range s.Slots {
		if slot == nil {
			continue
		}
		if slot.PlayerID != playerID {
			continue
		}
		// remove force = true after reconnection is implemented
		force = true
		slot.Connected = false
		slot.LastSeen = time.Now()
		if force {
			slot.PlayerID = 0
			slot.HeroUnitID = 0
		}
		log.Println("player",playerID,"removed from session", s.ID, "slot", i)
		break
	}
}

func (s *Session) SessionTick() {
	s.Tick++
	cmds := s.CommandQueue.Drain()
	s.ProcessCommands(cmds)
	s.World.Update()
	s.BroadcastState()
}

func (s *Session) ProcessCommands(
	cmds []command.Command,
) {
	for _, cmd := range cmds {
		switch c := cmd.(type) {
		case *command.MoveCommand:
			s.handleMove(c)
		case *command.AttackCommand:
			s.handleAttack(c)
		}
	}
}

func (s *Session) BroadcastState() {
    snapshot := protocol.WorldSnapshot{
        Tick: s.Tick,
    }
    for _, unit := range s.World.Units {
        snapshot.Units = append(
            snapshot.Units,
            protocol.UnitSnapshot{
                ID: uint64(unit.ID),
                X: unit.Position.X,
                Y: unit.Position.Y,
                HP: unit.HP,
            },
        )
    }    
    for _, client := range s.Clients {
        client.Send(snapshot)
    }
}

func (s *Session) CommandPush(cmd command.Command) {
    s.CommandQueue.Push(cmd)
}