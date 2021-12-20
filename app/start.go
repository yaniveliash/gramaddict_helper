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

		log.Println("pulling git")
		utils.PullGit()

		log.Println("starting ADB server")
		utils.SetADB()

		log.Println("unlocking device screen")
		utils.UnlockADB()

		message = message + "Starting app now...<br><br>"

		log.Printf("CHDIR %v", utils.Params.Home)

		args := []string{"-m", "run.py", "--config", string("accounts/"+utils.Params.Account) + "/config.yml"}

		log.Printf("Command 'python3 %v'", args)

		cmd := exec.Command("python3", args...)
		cmd.Dir = utils.Params.Home

		if err := cmd.Start(); err != nil {
			utils.HTMLMessage(201, c, "Not OK")
			return
		}

		go func() {
			err := cmd.Wait()
			utils.CheckError(err)
		}()

		if utils.IsRunning() {
			log.Println("Started")
			message = message + "Started<br>" + "<br><a href='/'>back</a><br>"
			utils.HTMLMessage(200, c, message)
		} else {
			log.Println("Failed to start")
			message = message + "Failed to start<br>"
			message = message + "<br><a href='/'>back</a><br>"
			utils.HTMLMessage(201, c, message)
			return
		}
	} else {
		log.Println("already running")
		message = message + "Already running...<br>" + "PID: " + utils.GetPID(utils.Params.Filter) + "<br><a href='/'>back</a><br>"
		utils.HTMLMessage(200, c, message)
	}

	log.Println("done")
}
