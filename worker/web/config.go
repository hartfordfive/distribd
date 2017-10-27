package web

import (
	"encoding/json"
	//"errors"

	"github.com/hartfordfive/distribd/worker"
)

type WebConfig struct {
	Locations     []Location `json:"locations"`      // List of locations to request
	RuntimePeriod int        `json:"runtime_period"` // Total time, in seconds, for which the test will run (if < 1, then runtime is infinite)
}

type Location struct {
	URL     string   `json:"url"`     // location to request
	Cookies []string `json:"cookies"` // List of cookies to send with the request
	Headers []string `json:"headers"` // List of headers to send with the request
	Weight  int      `json:"weight"`  // The weight/importance of this request in relation to all defined locations
}

func (wc *WebConfig) GetType() string {
	return "web"
}

func (wc *WebConfig) GetConfig() (*worker.ConfigMap, error) {

	var wcnf worker.ConfigMap
	data, err := json.Marshal(*wc)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &wcnf); err != nil {
		return nil, err
	}

	return &wcnf, nil
	//worker.ConfigMap
}
