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
		log.Println("I'm running, PID: ", utils.GetPID(utils.AppEnv.Filter))
		message = message + "Stopping app now...<br><br>"
		// log.Println(utils.GetPID(utils.AppEnv.Filter))
		cmd := exec.Command("kill", utils.GetPID(utils.AppEnv.Filter))
		err := cmd.Start()
		if err != nil {
			message = message + "Something went wrong while stopping...<br>"
			message = message + "<br><a href='/'>back</a><br>"
			utils.HtmlMessage(201, c, message)
			return
		}

		time.Sleep(2 * time.Second)
	}

	if !utils.IsRunning() {
		message = message + "Stopped app success.<br><br>"
		message = message + "<br><a href='/'>back</a><br>"
		log.Println("done")
		utils.HtmlMessage(200, c, message)
		return
	} else {
		message = message + "Failed to stop app.<br><br>"
		message = message + "<br><a href='/'>back</a><br>"
		utils.HtmlMessage(201, c, message)
		log.Println("Failed...")
		return
	}
}
