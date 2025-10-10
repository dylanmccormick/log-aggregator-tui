package log

import (
	"log/slog"
	"regexp"
	"time"
)

type StructuredLogParser struct{}

func (s *StructuredLogParser) CanParse(line string) bool {
	return true
}

func (s *StructuredLogParser) Parse(msg string) (LogMessage, error) {
	re := regexp.MustCompile(
		`^(?P<date>\d{4}-\d{2}-\d{2})` +
			`\s+` +
			`(?P<time>\d{2}:\d{2}:\d{2})` +
			`\s+` +
			`(?P<level>\w+)` +
			`(?:\s+\[(?P<component>[^\]]+)\])?` +
			`\s+` +
			`(?P<message>.+)$`)
	groups := namedGroups(re, msg)
	var lm LogMessage
	lm.Raw = msg

	layout := "2006-01-02 15:04:05"

	dateTime := groups["date"] + " " + groups["time"]

	timestamp, err := time.Parse(layout, dateTime)
	if err != nil {
		slog.Warn("Unable to parse timestamp. setting default", "input", dateTime)
		timestamp, _ = time.Parse(layout, "1776-07-04 11:11:11")
	}

	lm.Timestamp = timestamp
	lm.Level = groups["level"]
	lm.Message = groups["message"]
	lm.Component = groups["component"]

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
