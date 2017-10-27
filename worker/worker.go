package worker

import (
	"time"
)

type ConfigMap map[string]interface{}
type StatsMap map[string]interface{}

type WorkerConfig interface {
	GetType() string
	GetConfig() ConfigMap
}

type Worker interface {
	SetID(id string)
	GetID() string
	SetLastCheckIn(t *time.Time)
	GetStats() StatsMap
	GetConfig() ConfigMap
	String()
}
