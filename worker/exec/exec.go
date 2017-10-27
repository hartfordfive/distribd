package exec

import (
	"time"
)

type ExecNode struct {
	id          string
	lastCheckIn *time.Time
}

func (w *ExecNode) SetID(id string) {
	w.id = id
}

func (w *ExecNode) GetID() {
	return id
}

func (w *ExecNode) SetLastCheckIn(t *time.Time) {
	w.lastCheckIn = t
}

func (evt *ExecNode) GetStats() map[string]interface{} {
	return map[string]interface{}{}
}

func (evt *ExecNode) Print() {
	fmt.Printf("[NOTICE] %s (%d)\n", evt.message, evt.code)
}
