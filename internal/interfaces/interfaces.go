package interfaces

import (
	"sync"
	"time"
)

type (

	// Timestamp is a helper for (un)marhalling time
	Timestamp time.Time

	// HookMessage is the message we receive from Alertmanager
	HookMessage struct {
		Version           string            `json:"version"`
		GroupKey          string            `json:"groupKey"`
		Status            string            `json:"status"`
		Receiver          string            `json:"receiver"`
		GroupLabels       map[string]string `json:"groupLabels"`
		CommonLabels      map[string]string `json:"commonLabels"`
		CommonAnnotations map[string]string `json:"commonAnnotations"`
		ExternalURL       string            `json:"externalURL"`
		Alerts            []Alert           `json:"alerts"`
	}

	// Alert is a single alert.
	Alert struct {
		Labels      map[string]string `json:"labels"`
		Annotations map[string]string `json:"annotations"`
		StartsAt    string            `json:"startsAt,omitempty"`
		EndsAt      string            `json:"EndsAt,omitempty"`
	}

	// just an example alert store. in a real hook, you would do something useful
	alertStore struct {
		sync.Mutex
		capacity int
		alerts   []*HookMessage
	}
)
