package alertmanager

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
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

//Alertmanager -- Alertmanager
type Alertmanager struct {
	Host       string `yaml:"host"`
	URI        string `yaml:"uri"`
	Prefix     string `yaml:"prefix,omitempty"`
	mutex      sync.Mutex
	alertStore alertStore
}

//Config -- Config
func (alertmanager *Alertmanager) Config() *Alertmanager {
	return alertmanager
}

// Init -- Init
func (alertmanager *Alertmanager) Init() *Alertmanager {
	alertmanager.mutex.Lock()
	defer alertmanager.mutex.Unlock()
	alertmanager.alertStore = alertStore{}
	return alertmanager
}

// Handle -- Run
func (alertmanager *Alertmanager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	alertmanager.mutex.Lock()
	defer alertmanager.mutex.Unlock()
	switch r.Method {
	case http.MethodGet:
		alertmanager.getHandler(w, r)
	case http.MethodPost:
		alertmanager.postHandler(w, r)
	default:
		http.Error(w, "unsupported HTTP method", 400)
	}
}

func (alertmanager *Alertmanager) getHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	log.Println("Request")
	alertmanager.alertStore.Mutex.Lock()
	defer alertmanager.alertStore.Mutex.Unlock()

	if err := enc.Encode(alertmanager.alertStore.alerts); err != nil {
		log.Printf("error encoding messages: %v", err)
	}
}

func (alertmanager *Alertmanager) postHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var m HookMessage
	if err := dec.Decode(&m); err != nil {
		log.Printf("error decoding message: %v", err)
		http.Error(w, "invalid request body", 400)
		return
	}
	log.Println(m)
	alertmanager.alertStore.Mutex.Lock()
	defer alertmanager.alertStore.Mutex.Unlock()

	alertmanager.alertStore.alerts = append(alertmanager.alertStore.alerts, &m)
}

// Down -- Down
func (alertmanager *Alertmanager) Down() *Alertmanager {
	return alertmanager
}
