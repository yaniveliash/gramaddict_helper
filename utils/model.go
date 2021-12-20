package utils

import (
	"log"
	"os"
)

const (
	// ContentTypeHTML constant is the default output from the gin gonic router.
	ContentTypeHTML = "text/html; charset=utf-8"
)

// Env is a structure for the common parameters
type Env struct {
	Home       string
	AppLogPath string
	RunLogPath string
	Account    string
	Filter     string
}

// Params hold the environment parameters
var Params Env

// InitParams function initialized common bits.
func InitParams() {
	s, b := os.LookupEnv("GH")
	if b {
		Params.Home = s
	} else {
		Params.Home = os.Getenv("HOME") + "/gramaddict"
	}

	s, b = os.LookupEnv("ACC")
	if b {
		Params.Account = s
	} else {
		envError("ACC env var must be set")
	}

	s, b = os.LookupEnv("FILTER")
	if b {
		Params.Filter = s
	} else {
		Params.Filter = "run.py"
	}

	s, b = os.LookupEnv("LOG")
	if b {
		Params.AppLogPath = Params.Home + "/" + s
	} else {
		Params.AppLogPath = Params.Home + "/logs/" + Params.Account + ".log"
	}
}

func envError(message string) {
	log.Fatal(message)
}
