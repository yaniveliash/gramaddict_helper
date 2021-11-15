package utils

import (
	"log"
	"os"
)

const (
	ContentTypeHTML = "text/html; charset=utf-8"
)

type Env struct {
	Home       string
	AppLogPath string
	RunLogPath string
	Account    string
	Filter     string
}

var (
	AppEnv Env
)

func InitEnv() {
	s, b := os.LookupEnv("GH")
	if b {
		AppEnv.Home = s
	} else {
		AppEnv.Home = os.Getenv("HOME") + "/gramaddict"
	}

	s, b = os.LookupEnv("ACC")
	if b {
		AppEnv.Account = s
	} else {
		envError("ACC env var much be set")
	}

	s, b = os.LookupEnv("FILTER")
	if b {
		AppEnv.Filter = s
	} else {
		AppEnv.Filter = "run.py"
	}

	s, b = os.LookupEnv("LOG")
	if b {
		AppEnv.AppLogPath = AppEnv.Home + "/" + s
	} else {
		AppEnv.AppLogPath = AppEnv.Home + "/logs/" + AppEnv.Account + ".log"
	}
}

func envError(message string) {
	log.Fatal(message)
}
