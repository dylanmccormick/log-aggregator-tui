package log

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"time"
)

type LogParser interface {
	CanParse(line string) bool
	Parse(line string) (LogMessage, error)
	Priority() int
	Name() string
}

type ParserRegistry struct {
	Parsers []LogParser
}

type LogMessage struct {
	Raw       string
	Timestamp time.Time
	Level     string
	Message   string
	Component string
}

func NewParserRegistry() *ParserRegistry {
	return &ParserRegistry{
		[]LogParser{
			NewStructuredLogParser(),
		},
	}
}

func (pr *ParserRegistry) Parse(line string) (LogMessage, error) {
	for _, parser := range pr.Parsers {
		if parser.CanParse(line) {
			return parser.Parse(line)
		}
	}
	return LogMessage{}, fmt.Errorf("no parser available for the given line format")
}

func ReadLogFile(filePath string) ([]LogMessage, error) {
	var messages []LogMessage
	file, err := os.Open(filePath)
	if err != nil {
		slog.Error("Unexpected error with opening file", "filePath", filePath, "error", err)
		return []LogMessage{}, fmt.Errorf("failed to open file%s: %w", filePath, err)
	}
	defer file.Close()
	slog.Info("Reading log file", "filePath", filePath)

	pr := NewParserRegistry()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		parsedLog, err := pr.Parse(line)
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

func namedGroups(re *regexp.Regexp, text string) map[string]string {
	matches := re.FindStringSubmatch(text)
	if matches == nil {
		return nil
	}
	results := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			results[name] = matches[i]
		}
	}
	return results
}

// func parseLog(msg string) (LogMessage, error) {
// 	re := regexp.MustCompile(
// 		`^(?P<date>\d{4}-\d{2}-\d{2})` +
// 			`\s+` +
// 			`(?P<time>\d{2}:\d{2}:\d{2})` +
// 			`\s+` +
// 			`(?P<level>\w+)` +
// 			`(?:\s+\[(?P<component>[^\]]+)\])?` +
// 			`\s+` +
// 			`(?P<message>.+)$`)
// 	groups := namedGroups(re, msg)
// 	var lm LogMessage
// 	lm.Raw = msg
//
// 	layout := "2006-01-02 15:04:05"
//
// 	dateTime := groups["date"] + " " + groups["time"]
//
// 	timestamp, err := time.Parse(layout, dateTime)
// 	if err != nil {
// 		slog.Warn("Unable to parse timestamp. setting default", "input", dateTime)
// 		timestamp, _ = time.Parse(layout, "1776-07-04 11:11:11")
// 	}
//
// 	lm.Timestamp = timestamp
// 	lm.Level = groups["level"]
// 	lm.Message = groups["message"]
// 	lm.Component = groups["component"]
//
// 	return lm, nil
// }
