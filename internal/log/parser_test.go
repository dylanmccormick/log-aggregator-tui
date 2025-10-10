package log

import "testing"

func TestParseLog(t *testing.T) {
	tests := []struct {
		input             string
		expectedLevel     string
		expectedComponent string
		expectedMessage   string
	}{
		{
		 "2023-10-07 14:32:15 INFO [user-service] User login successful: user_id=12345",
		 "INFO",
		 "user-service",
		 "User login successful: user_id=12345",
		},
		{
		 "2023-10-07 14:32:15 INFO User login successful: user_id=12345",
		 "INFO",
		 "",
		 "User login successful: user_id=12345",
		},

	}
	pr := NewParserRegistry()

	for i, test := range(tests) {
		logMsg, err := pr.Parse(test.input)
		if err != nil {
			t.Errorf("failed parsing log message with error: %s. test_no:%d", err, i)
		}
		if logMsg.Raw != test.input {
			t.Errorf("Incorrect raw log message. expected=%s got=%s test_no:%d", test.input, logMsg.Raw, i)
		}
		if logMsg.Level != test.expectedLevel{
			t.Errorf("Incorrect log level. expected=%s got=%s test_no:%d", test.expectedLevel, logMsg.Level, i)
		}
		if logMsg.Component != test.expectedComponent{
			t.Errorf("Incorrect log component. expected=%s got=%s test_no:%d", test.expectedComponent, logMsg.Component, i)
		}
		if logMsg.Message != test.expectedMessage{
			t.Errorf("Incorrect log message. expected=%s got=%s test_no:%d", test.expectedMessage, logMsg.Message, i)
		}
	}
}
