package logging

import (
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
)

// Logger -- Logger
type Logger struct {
	LogFile string `yaml:"logFile"`
	Debug   bool   `yaml:"debug"`
}

//SetupLogging -- SetupLogging
func SetupLogging(cfg *Logger) {

	if cfg.LogFile != "" {
		logfile, err := os.Open(cfg.LogFile)
		if err != nil {
			log.Fatalf("Unable to read config: %+v", err)
		}
		log.SetOutput(logfile)
	}
	if cfg.Debug == true {
		log.SetLevel(log.DebugLevel)
	}
	log.SetLevel(log.InfoLevel)
	log.Printf("%+v", cfg)
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

//CheckError -- CheckError
func CheckError(msg string, err error) {
	if err != nil {
		if stackErr, ok := err.(stackTracer); ok {
			log.WithField("stacktrace", fmt.Sprintf("%+v", stackErr.StackTrace())).Errorf("Error: %s %+v", msg, err)
		} else {
			log.Errorf("Error: %s %+v", msg, err)
		}
	}
}

//Log -- log
func Log(msg string) {
	log.Println(msg)
}

//Logf -- logf
func Logf(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}
