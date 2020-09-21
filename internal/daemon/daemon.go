package daemon

import (
	//"github.com/m1keru/pushbot/internal/alertmanager"
	"fmt"
	"github.com/m1keru/pushbot/internal/config"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

//SpinUp -- SpinUp
func SpinUp(cfg *config.Config, wg *sync.WaitGroup) error {
	alertmanager := &cfg.Alertmanager
	http.Handle(cfg.Alertmanager.URI, alertmanager)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Alertmanager.Host, cfg.Daemon.Port), nil))
	return nil
}
