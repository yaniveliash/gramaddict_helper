package app

import (
	"log"
	"os/exec"
	"time"

	"github.com/geostant/gramaddict_helper/utils"
	"github.com/gin-gonic/gin"
)

// StopApp function stop the running process and clean the env.
func StopApp(c *gin.Context) {
	var message string
	if utils.IsRunning() {
		log.Println("I'm running, PID: ", utils.GetPID(utils.Params.Filter))
		message = message + "Stopping app now...<br><br>"
		// log.Println(utils.GetPID(utils.Params.Filter))
		cmd := exec.Command("kill", utils.GetPID(utils.Params.Filter))
		err := cmd.Start()
		if err != nil {
			message = message + "Something went wrong while stopping...<br>"
			message = message + "<br><a href='/'>back</a><br>"
			utils.HTMLMessage(201, c, message)
		}

		time.Sleep(2 * time.Second)

		if !utils.IsRunning() {
			message = message + "Stopped app success.<br><br>"
			message = message + "<br><a href='/'>back</a><br>"
			log.Println("done")
			utils.HTMLMessage(200, c, message)
		} else {
			message = message + "Failed to stop app.<br><br>"
			message = message + "<br><a href='/'>back</a><br>"
			utils.HTMLMessage(201, c, message)
			log.Println("Failed...")
		}
	} else {
		message = message + "APP isn't running....<br><br>"
		message = message + "<br><a href='/'>back</a><br>"
		log.Println("done")
		utils.HTMLMessage(200, c, message)
	}
}
