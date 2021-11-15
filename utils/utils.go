package utils

import (
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	return r
}

// HtmlMessage function responsible for the output message from the router.
func HtmlMessage(statusCode int, c *gin.Context, message string) {
	c.Data(statusCode, ContentTypeHTML, []byte(message))
}

// IsRunning check the process status
func IsRunning() bool {
	return GetPID(AppEnv.Filter) != ""
}

// GetPIDByFile function return the PID number
// func GetPIDByFile() string {
// 	rawData, err := os.ReadFile(AppEnv.PathToPid)
// 	CheckError(err)
// 	return string(rawData)
// }

// GetPID function return the PID number
func GetPID(filter string) string {
	cmd := exec.Command("ps", "aux")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range strings.Split(string(out), "\n") {
		if strings.Contains(l, filter) {
			f := strings.Fields(l)
			return f[1]
		}
	}
	return ""
}

func getAppLog(c *gin.Context, appLogPath string) string {
	vals := c.Request.URL.Query()

	lines, err := strconv.Atoi(vals.Get("lines"))
	if err != nil {
		lines = 41
	}

	data, err := ioutil.ReadFile(appLogPath)
	CheckError(err)

	s := strings.Split(string(data), "\n")
	var newS string
	for i := len(s) - 1; i > len(s)-lines; i-- {
		if i == 0 {
			break
		}

		newS = newS + s[i] + "<br>"
	}

	return newS
}

// GetStatus responds with the list of all albums as JSON.
func GetStatus(c *gin.Context) {
	dt := time.Now()

	if IsRunning() {
		c.Data(200, ContentTypeHTML, []byte("<html>["+dt.String()+"] ✅ <br><a href='/stop'>stop app</a><h3>app log</h3><code>"+getAppLog(c, AppEnv.AppLogPath)+"</code></html>"))
	} else {
		c.Data(200, ContentTypeHTML, []byte("<html>["+dt.String()+"] ❌ <br><a href='/start'>start app</a><h3>app log</h3><code>"+getAppLog(c, AppEnv.AppLogPath)+"</code></html>"))
	}
}

// CheckError function handle errors
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// PullGit function is...
func PullGit() {
	cmd := exec.Command("git", "pull")
	cmd.Dir = AppEnv.Home
	out, err := cmd.Output()

	if err != nil {
		log.Println(out)
		log.Fatal(err)
	}
	log.Println(string(out))
}

func SetADB() {
	cmd := exec.Command("adb", "devices")
	out, err := cmd.Output()

	if err != nil {
		log.Println(out)
		log.Fatal(err)
	}
	log.Println(string(out))
}

func UnlockADB() {
	cmd := exec.Command("adb", "shell", "input", "keyevent", "82")
	out, err := cmd.Output()

	if err != nil {
		log.Println(out)
		log.Fatal(err)
	}

	cmd = exec.Command("adb", "shell", "input", "swipe", "100", "500", "100", "1450", "100")
	out, err = cmd.Output()

	if err != nil {
		log.Println(out)
		log.Fatal(err)
	}
}
