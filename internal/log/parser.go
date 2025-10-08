package log

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"
)

type logMessage struct {
	Raw       string
	Timestamp time.Time
	Level     string
	Message   string
	Component string
}

func ReadLogFile(filePath string) ([]logMessage, error){
	var messages []logMessage
	file, err := os.Open(filePath)
	if err != nil {
		slog.Error("Unexpected error with opening file", "filePath", filePath, "error", err)
		return []logMessage{}, fmt.Errorf("failed to open file%s: %w", filePath, err)
	}
	defer file.Close()
	slog.Info("Reading log file", "filePath", filePath)

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		parsedLog, err := parseLog(line)
		if err != nil {
			slog.Error("Unable to parse log", "LineNumber", lineNumber, "RawLog", line, "error", err)
			continue
		}
		messages = append(messages, parsedLog)
	}

	if err := scanner.Err(); err != nil {
		return messages, fmt.Errorf("error reading file: %w", err)
	}

	return messages, nil
}

func parseLog(msg string) (logMessage, error) {
	var lm logMessage
	lm.Raw = msg
	strings := strings.SplitN(msg, " ", 5)
	date := strings[0]
	t := strings[1]
	level := strings[2]
	component := strings[3]
	message := strings[4]

	layout := "2006-01-02 15:04:05"

	timestamp, err := time.Parse(layout, date+" "+t)
	if err != nil {
		slog.Warn("Unable to parse timestamp. setting default", "input", date+" "+t)
		timestamp, _ = time.Parse(layout, "1776-07-04 11:11:11")
	}

	lm.Timestamp = timestamp
	lm.Level = level
	lm.Message = message
	lm.Component = cleanString(component)

	return lm, nil
}

func cleanString(field string) string {
	if field[0] == '[' {
		field = field[1:]
	}

	if field[len(field)-1] == ']' {
		field = field[:len(field)-1]
	}

	return field
}
