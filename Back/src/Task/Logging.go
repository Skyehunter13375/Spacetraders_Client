package Task

import "log"
import "os"
import "fmt"
import "time"

func LogErr(msg string) {
	CFG, _     := GetConfig()
	logFile, _ := os.OpenFile(CFG.LOG.ErrPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(logFile)
	log.Printf("%s", msg)
	logFile.Close()

	currentTime := time.Now().UTC()
	formatdTime := currentTime.Format("01/02/2006 15:04:05")
	fmt.Println(formatdTime + " - " + msg)
}

func LogActivity(msg string) {
	CFG, _     := GetConfig()
	logFile, _ := os.OpenFile(CFG.LOG.ActPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(logFile)
	log.Printf("%s", msg)
	logFile.Close()

	currentTime := time.Now().UTC()
	formatdTime := currentTime.Format("01/02/2006 15:04:05")
	fmt.Println(formatdTime + " - " + msg)
}


