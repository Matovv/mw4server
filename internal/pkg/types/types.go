package types

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
