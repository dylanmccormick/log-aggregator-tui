package log

import "time"

type logMessage struct {
	Raw       string
	Timestamp time.Time
	Level     string
	Message   string
	Component string
}

func readLogFile(filePath string) []logMessage {
	var messages []logMessage

	return messages
}
