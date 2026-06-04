package network

type JoinSessionResponse struct {
    PlayerID  uint64   `json:"player_id"`
    SlotIndex uint8    `json:"slot_index"`
    UnitID    uint64   `json:"unit_id"`
    Tick      uint64   `json:"tick"`     
}