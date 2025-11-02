package General

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func LogErr(msg string) {
	logFile, _ := os.OpenFile("Logs/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(logFile)
	log.Printf("%s", msg)
	logFile.Close()
}

func LogActivity(msg string) {
	logFile, _ := os.OpenFile("Logs/activity.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(logFile)
	log.Printf("%s", msg)
	logFile.Close()
}

func ProgressBar(curr int, max int) string {
	var bar strings.Builder
	var percent int
	totalWidth := 20
	if max == 0 {
		percent = 100
	} else {
		percent = (curr / max) * 100
	}
	filled := (percent * totalWidth) / 100
	empty := totalWidth - filled

	fmt.Fprintf(&bar,
		"[[green]%s[-][black]%s[-]] %d%%",
		strings.Repeat("■", filled),
		strings.Repeat(" ", empty),
		percent,
	)

	return bar.String()
}
