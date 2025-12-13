package Task

import "log"
import "os"

func LogErr(msg string) {
	CFG, _     := GetConfig()
	logFile, _ := os.OpenFile(CFG.LOG.ErrPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(logFile)
	log.Printf("%s", msg)
	logFile.Close()
}

func LogActivity(msg string) {
	CFG, _     := GetConfig()
	logFile, _ := os.OpenFile(CFG.LOG.ActPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(logFile)
	log.Printf("%s", msg)
	logFile.Close()
}


