package General

import "fmt"
import "log"
import "os"
import "strings"

func LogErr(msg string) {
	logFile, _ := os.OpenFile("/var/log/spacetraders/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(logFile)
	log.Printf("%s", msg)
	logFile.Close()
}

func LogActivity(msg string) {
	logFile, _ := os.OpenFile("/var/log/spacetraders/activity.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(logFile)
	log.Printf("%s", msg)
	logFile.Close()
}

func ProgressBar(curr, req, max int) string {
	if max <= 0 { return "N/A" }

	var bar strings.Builder
	totalWidth := 20

	// Clamp values so we don't go below 0 or above max
	if curr < 0   { curr = 0 }
	if curr > max { curr = max }
	if req  < 0   { req = 0 }
	if req  > max { req = max }

	// Compute percentages
	percent := int(float64(curr) / float64(max) * 100)
	filled := (percent * totalWidth) / 100
	reqPos := (int(float64(req) / float64(max) * float64(totalWidth)))

	// Build the bar
	for i := range totalWidth {
		switch {
			case i < filled:
				bar.WriteString("■")
			case i == reqPos:
				bar.WriteString("|") // marker for required threshold
			default:
				bar.WriteString(" ")
			}
	}

	return fmt.Sprintf("[[green]%s[-]] %d%%", bar.String(), percent)
}
