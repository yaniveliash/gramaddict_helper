package app

import (
	"log"
	"os/exec"

	"github.com/geostant/gramaddict_helper/utils"
	"github.com/gin-gonic/gin"
)

// StartApp function start the process and set PID.
func StartApp(c *gin.Context) {
	var message string
	if !utils.IsRunning() {
		log.Println("not running")

		utils.PullGit()
		utils.SetADB()
		utils.UnlockADB()

		message = message + "Starting app now...<br><br>"

		args := []string{"-m", "run.py", "--config", string("accounts/" + utils.AppEnv.Account), "/config.yml"}
		cmd := exec.Command("python3", args...)
		cmd.Dir = utils.AppEnv.Home
		err := cmd.Start()

		if err != nil {
			log.Println("something wrong")
			message = message + "Something went wrong.<br><br>"
			message = message + "<br><a href='/'>back</a><br>"
			utils.HtmlMessage(200, c, message)
			return
		}

		if utils.IsRunning() {
			message = message + "Started<br>" + "<br><a href='/'>back</a><br>"
			utils.HtmlMessage(200, c, message)
		} else {
			message = message + "Failed to start<br>"
			message = message + "<br><a href='/'>back</a><br>"
			utils.HtmlMessage(201, c, message)
			return
		}
	} else {
		log.Println("already running")
		message = message + "Already running...<br>" + "PID: " + utils.GetPID(utils.AppEnv.Filter) + "<br><a href='/'>back</a><br>"
		utils.HtmlMessage(200, c, message)
	}
	log.Println("done")
}
