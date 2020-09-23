package main

import (
	"flag"
	"github.com/m1keru/pushbot/internal/config"
	"github.com/m1keru/pushbot/internal/daemon"
	"github.com/m1keru/pushbot/internal/logging"
	log "github.com/sirupsen/logrus"
	"sync"
)

func main() {
	configpath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()
	var cfg config.Config
	if err := cfg.Setup(configpath); err != nil {
		log.Fatalf("%+v", err)
	}

	logging.SetupLogging(&cfg.Logger)
	var wg sync.WaitGroup
	//wg.Add(1)
	daemon.SpinUp(&cfg, &wg)

	/*inPipe := make(chan []byte)
	outPipe := make(chan string)
	go telegram.Run(&cfg, &wg, &audioPipe, &textPipe)
	go speech.Run(&cfg, &wg, &audioPipe, &textPipe)
	log.Debug("waiting on WaitGroup")
	wg.Wait()
	*/

}
