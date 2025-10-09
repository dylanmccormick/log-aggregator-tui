package log

import (
	"log/slog"
	"strings"
	"time"
)

type StructuredLogParser struct{}

func (s *StructuredLogParser) CanParse(line string) bool {
	return true
}

func (s *StructuredLogParser) Parse(msg string) (LogMessage, error) {
	var lm LogMessage
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

func (s *StructuredLogParser) Priority() int {
	return 100
}

func (s *StructuredLogParser) Name() string {
	return "StructuredLogParser"
}

func NewStructuredLogParser() LogParser {
	
	return &StructuredLogParser{}
}
