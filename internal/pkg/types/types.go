package types

const (
	WorldScale = 1000
)

type SessionID string
type PlayerID uint64
type UnitID uint64
type ObjectID uint64
type PrefabID uint64

type TargetID uint64
func (tid TargetID) ToUnitID() UnitID {
	return UnitID(tid)
}
func (tid TargetID) ToObjectID() ObjectID {
	return ObjectID(tid)
}
type Vec2 struct {
	X int64
	Y int64
}
type Rotation uint16
type Transform struct {
	Position Vec2
	Rotation
}
func (t *Transform) Set(newT Transform) {
	t.Position.X = newT.Position.X
	t.Position.Y = newT.Position.Y
	rot := newT.Rotation
	if rot > 359 {
		rot = 0
	}
	t.Rotation = rot
}
func NewTransform(x, y int64, rotation int16) *Transform{
	newPos := Vec2{
		X: x,
		Y: y,
	}
	return &Transform{
		Position: newPos,
		Rotation: Rotation(rotation),
	}
}

type Faction string
const (
	FactionNone Faction = "NONE"
	FactionNeutralGood Faction = "NEUTRAL_GOOD"
	FactionNeutralEvil Faction = "NEUTRAL_EVIL"
	FactionA Faction = "A"
	FactionB Faction = "B"
	FactionC Faction = "C"
	FactionD Faction = "D"
)
