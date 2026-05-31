package types

type SessionID string
type PlayerID uint64
type UnitID uint64
type ObjectID uint64

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
