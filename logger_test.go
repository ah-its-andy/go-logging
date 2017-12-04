package logging

import "testing"

func Test_Logger(t *testing.T) {
	logger := NewLogger()
	logger.TryAddProvider(NewConsoleProvider())
	logger.Errorln("testing")
}
