package endpoints

import (
	"encoding/json"
	"github.com/m1keru/pushbot/internal/databases"
	"github.com/m1keru/pushbot/internal/interfaces"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

//Alertmanager -- Alertmanager
type Alertmanager struct {
	Host   string `yaml:"host"`
	URI    string `yaml:"uri"`
	Prefix string `yaml:"prefix,omitempty"`
	Mutex  sync.Mutex
	Dblink *databases.RabbitMQ
}

// Handle -- Run
func (alertmanager *Alertmanager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, "MethodGet is not allowed", http.StatusMethodNotAllowed)
}

func (alertmanager *Alertmanager) postHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var m interfaces.HookMessage
	if err := dec.Decode(&m); err != nil {
		log.Printf("error decoding message: %+v \nerror: %v\n", m, err)
		http.Error(w, "invalid request body", 400)
		return
	}
	log.Printf("Alertmanager Publish %+v %+v", m, alertmanager.Dblink)
	alertmanager.Mutex.Lock()
	alertmanager.Dblink.Init()
	alertmanager.Dblink.Publish(m)
	alertmanager.Mutex.Unlock()
	//alertmanager.alertStore.alerts = append(alertmanager.alertStore.alerts, &m)
}

// Down -- Down
func (alertmanager *Alertmanager) Down() *Alertmanager {
	return alertmanager
}
