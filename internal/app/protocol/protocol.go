package protocol

type ObjectSnapshot struct {
    PrefabID    uint64
	ID          uint64
	X           int64
	Y           int64
}

type UnitSnapshot struct {
    GameModelID string
    Size        uint64
	ID          uint64
	Faction     string
	X           int64
	Y           int64
	Hp          uint64
	Mana        uint64
}

type WorldSnapshot struct {
	Tick    uint64
	Units   []UnitSnapshot
	Objects []ObjectSnapshot
}
