package log

import "testing"


func TestParseLog(t *testing.T) {
	msg := "2023-10-07 14:32:15 INFO [user-service] User login successful: user_id=12345"
	logMsg, err := parseLog(msg)
	if err != nil {
		t.Errorf("failed parsing log message with error: %s", err)
	}
	if logMsg.Raw != msg {
		t.Errorf("Incorrect raw log message. expected=%s got=%s", msg, logMsg.Raw)
	}
	if logMsg.Level != "INFO" {
		t.Errorf("Incorrect log level. expected=%s got=%s", "INFO", logMsg.Level)
	}
	if logMsg.Component != "user-service" {
		t.Errorf("Incorrect log component. expected=%s got=%s", "user-service", logMsg.Component)
	}
}

