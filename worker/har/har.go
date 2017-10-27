package http

import (
	"time"

	"github.com/satori/go.uuid"
)

type HarNode struct {
	id          string
	lastCheckIn *time.Time
	taskChanSize int
	tasks       chan string // channel of json encoded payloads
}

/*
	New
*/
func NewHarNode(taskQueueSize int) *HarNode {
	return &WebNode{
		id:          uuid.NewV4().String(),
		lastCheckIn: nil,
		taskChanSize: taskQueueSize,
	}
}

func (w *HarNode) init()
	w.tasks = make(chan string, taskQueueSize)
}

func (w *HarNode) SetID(id string) {
	w.id = id
}

func (w *HarNode) GetID() {
	return id
}

func (w *HarNode) SetLastCheckIn(t *time.Time) {
	w.lastCheckIn = t
}

func (evt *HarNode) GetStats() map[string]interface{} {
	return map[string]interface{}{}
}

func (evt *HarNode) Print() {
	fmt.Printf("[NOTICE] %s (%d)\n", evt.message, evt.code)
}
