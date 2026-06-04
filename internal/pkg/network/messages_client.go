package network

type JoinSessionRequest struct {
    SessionID string `json:"session_id"`
}

type MoveRequest struct {
	UnitID uint64  `json:"unit_id"`
	X      int64   `json:"x"`
	Y      int64   `json:"y"`
}

type AttackRequest struct {
    UnitID   uint64   `json:"unit_id"`
    TargetID uint64   `json:"target_id"`
}