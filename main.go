package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/geostant/gramaddict_helper/app"
	"github.com/geostant/gramaddict_helper/utils"
)

func main() {
	helpFlag := flag.Bool("help", false, "print out the help")
	flag.Parse()

	if *helpFlag {
		help()
	}

	utils.InitParams()

	router := utils.SetupRouter()
	router.GET("/", utils.GetStatus)
	router.GET("/start", app.StartApp)
	router.GET("/stop", app.StopApp)

	err := router.Run(":8080")
	utils.CheckError(err)
}

func help() {
	fmt.Println(`
Environment variables

=========================================================================================================================

ACC [Mandaotory]              Instagram account name

GH [Optional]                 Gramaddict installation directory (default to '$HOME/gramaddict')

FILTER [Optional]             Used to determine if the process running by filtering the processes (default to 'run.py')

LOG [Optional]                Path to gramaddict account log file (default to '$HOME/gramaddict/logs/$INSTAGRAM_ACCOUNT.log'
 `)
	os.Exit(0)
}
