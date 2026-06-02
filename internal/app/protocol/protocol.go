package protocol

type ObjectSnapshot struct {
	ID uint64
	X  int64
	Y  int64
}

type UnitSnapshot struct {
	ID      uint64
	Faction string
	X       int64
	Y       int64
	HP      int64
	Mana    int64
}

type WorldSnapshot struct {
	Tick    uint64
	Units   []UnitSnapshot
	Objects []ObjectSnapshot
}
