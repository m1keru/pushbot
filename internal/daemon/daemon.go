package daemon

import (
	"fmt"
	"github.com/m1keru/pushbot/internal/config"
	"github.com/m1keru/pushbot/internal/logging"
	"net/http"
	"sync"
)

//SpinUp -- SpinUp
func SpinUp(cfg *config.Config, wg *sync.WaitGroup) error {
	alertmanager := &cfg.Alertmanager
	alertmanager.Dblink = &cfg.Rabbitmq
	http.Handle(cfg.Alertmanager.URI, alertmanager)
	logging.CheckError("Daemon failed", http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Alertmanager.Host, cfg.Daemon.Port), nil))
	return nil
}
