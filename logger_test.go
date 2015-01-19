package logger

import "testing"

var base = Logger{
	name:      "testLog",
	stdout:    true,
	file:      true,
	raw:       false,
	verbosity: 2,
	path:      "logs/",
	tlayout:   "2006-01-02 15:04:05 MST",
}

func TestNew(t *testing.T) {
	testLog := *New("testLog")

	if !compareLoggers(base, testLog) {
		t.Error("Expected ", base, " Got ", testLog)
	}
	if len(loggers) != 1 {
		t.Error("Expected 1 Got ", len(loggers))
	}
}

func TestGet(t *testing.T) {
	testLog := *Get("testLog")

	if !compareLoggers(base, testLog) {
		t.Error("Expected ", base, " Got ", testLog)
	}
}

func TestGlobalVerbose(t *testing.T) {
	Verbose(3)
	if GetVerboseLevel() != 3 {
		t.Error("Expected 3 Got ", GetVerboseLevel())
	}
	Verbose(5)
	if GetVerboseLevel() != 3 {
		t.Error("Expected 3 Got ", GetVerboseLevel())
	}
	Verbose(-2)
	if GetVerboseLevel() != 0 {
		t.Error("Expected 0 Got ", GetVerboseLevel())
	}
}

func TestLoggerVerbose(t *testing.T) {
	testLog := Get("testLog")
	if testLog.verbosity != 2 {
		t.Error("Expected 2 Got ", testLog.verbosity)
	}

	testLog.Verbose(3)
	if testLog.verbosity != 3 {
		t.Error("Expected 3 Got ", testLog.verbosity)
	}

	testLog.Verbose(6)
	if testLog.verbosity != 3 {
		t.Error("Expected 3 Got ", testLog.verbosity)
	}

	testLog.Verbose(-25)
	if testLog.verbosity != 0 {
		t.Error("Expected 3 Got ", testLog.verbosity)
	}
}

func TestLoggerPath(t *testing.T) {
	// Default path
	testLog := Get("testLog")
	if testLog.path != "logs/" {
		t.Error("Expected 'logs/' Got ", testLog.path)
	}

	// Path with no trailing forward slash
	testLog.Path("test")
	if testLog.path != "test/" {
		t.Error("Expected 'test/' Got ", testLog.path)
	}

	// Path with trailing forward slash
	testLog.Path("loggers/")
	if testLog.path != "loggers/" {
		t.Error("Expected 'loggers/' Got ", testLog.path)
	}
}

func compareLoggers(a, b Logger) bool {
	if &a == &b {
		return true
	}
	if a.name != b.name || a.path != b.path || a.tlayout != b.tlayout {
		return false
	}
	if a.stdout != b.stdout || a.file != b.file || a.raw != b.raw {
		return false
	}
	if a.verbosity != b.verbosity {
		return false
	}
	if !a.t.start.Equal(b.t.start) {
		return false
	}
	if a.t.running != b.t.running {
		return false
	}
	return true
}
