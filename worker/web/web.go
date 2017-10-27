package web

import (
	"fmt"
	"time"

	"github.com/satori/go.uuid"
	//"github.com/coreos/etcd/embed"
)

type WebNode struct {
	id           string
	lastCheckIn  *time.Time
	taskChanSize int
	tasks        chan string // channel of json encoded payloads
}

func init() {

}

func NewWebNode(taskQueueSize int) *WebNode {
	wn := &WebNode{
		id:           uuid.NewV4().String(),
		lastCheckIn:  nil,
		taskChanSize: taskQueueSize,
	}
	wn.tasks = make(chan string, taskQueueSize)
	return wn
}

func (w *WebNode) init() {

}

func (w *WebNode) SetID(id string) {
	w.id = id
}

func (w *WebNode) GetID() string {
	return w.id
}

func (w *WebNode) SetLastCheckIn(t *time.Time) {
	w.lastCheckIn = t
}

func (w *WebNode) GetLastCheckIn(t *time.Time) {
	return w.lastCheckIn
}

func (w *WebNode) GetStats() map[string]interface{} {
	return map[string]interface{}{}
}

func (w *WebNode) Print() {
	fmt.Printf("[%s] %s (%d)\n", w.GetID(), w.GetLastCheckIn())
}
