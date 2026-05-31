package network

import (
	"log"
	"net/http"
	"sync/atomic"

	"github.com/Matovv/mw4server/internal/pkg/types"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
    SessionManager sessionManager
    NextPlayerID atomic.Uint64
}

func NewServer(
	sm sessionManager,
) *Server {
	return &Server{
		SessionManager: sm,
	}
}

func (s *Server) Start(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc(
		"/ws",
		s.handleWebsocket,
	)
	log.Println("server online!")
	return http.ListenAndServe(
		addr,
		mux,
	)
}

func (s *Server) handleWebsocket(
	w http.ResponseWriter,
	r *http.Request,
) {
	conn, err := upgrader.Upgrade(
		w,
		r,
		nil,
	)
	if err != nil {
		return
	}
	s.HandleConnection(conn)
}

func (s *Server) HandleConnection(
	conn *websocket.Conn,
) {
	playerID := s.NextPlayerID.Add(1)
	client := &Client{
		ID:       playerID,
		Conn:     conn,
		SendChan: make(chan []byte, 128),
	}
	log.Println("new connection - id:", playerID)
	go client.writeLoop()
	client.readLoop(s)
	s.HandleDisconnect(client, false)
}

func (s *Server) HandleDisconnect(
	client *Client,
	force bool,
) {
	if client.SessionID == "" {
		return
	}
	session, err := s.SessionManager.GetSession(
		client.SessionID,
	)
	if err != nil {
		log.Println("error on player disconnect:", err.Error())
		return
	}
	session.RemovePlayer(types.PlayerID(client.ID), force)
	client.Close()
	log.Println(
		"player", client.ID, "disconnected:",
	)
}

func (s *Server) HandlePacket(
    client *Client,
    packet Packet,
) {
    switch packet.Type {
    case "join_session":
		log.Println("handling join_session")
		//log.Println("data:", string(packet.Data))
        err := s.handleJoinSession(
            client,
            packet.Data,
        )
		if err != nil {
			log.Println("failed join_session:", err)			
		}
    case "move":
		log.Println("handling move")
        err := s.handleMove(
            client,
            packet.Data,
        )
		if err != nil {
			log.Println("failed move:", err)
		}
    }
}