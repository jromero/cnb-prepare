package testhelpers

import (
	"testing"

	"github.com/jromero/cnb-prepare/pkg/preparer"
)

func NewLogger(t *testing.T) preparer.Logger {
	return &testLogger{t: t}
}

type testLogger struct {
	t *testing.T
}

func (l *testLogger) Debug(format string, args ...interface{}) {
	l.t.Helper()
	l.t.Logf("[DEBUG] "+format, args...)
}

func (l *testLogger) Info(format string, args ...interface{}) {
	l.t.Helper()
	l.t.Logf("[INFO] "+format, args...)
}

func (l *testLogger) Warn(format string, args ...interface{}) {
	l.t.Helper()
	l.t.Logf("[WARN] "+format, args...)
}
