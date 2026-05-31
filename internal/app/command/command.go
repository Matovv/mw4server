package command

type Type uint8

const (
	MoveCommandType Type = iota
	AttackCommandType
	CastSpellCommandType
)

type Command interface {
    CommandType() Type
}