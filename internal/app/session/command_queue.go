package session

import (
	"sync"

	"github.com/Matovv/mw4server/internal/app/command"
)

type CommandQueue struct {
	mu sync.Mutex
	commands []command.Command
}

func (q *CommandQueue) Push(cmd command.Command) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.commands = append(q.commands, cmd)
}

func (q *CommandQueue) Drain() []command.Command {
	q.mu.Lock()
	defer q.mu.Unlock()
	cmds := q.commands
	q.commands = nil
	return cmds
}
