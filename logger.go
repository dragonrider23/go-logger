package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Colors that can be used for Error()
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	White   = "\033[37m"
	Grey    = "\x1B[90m"
)

// Type logger is the struct returned and used for logging
// The user can set its name, logfile location, time layout,
// and if it's shown in stdout.
type logger struct {
	name, location, tlayout string
	stdout                  bool
}

// Loggers is the collection of all logger types used.
// A logger can be called by their names instead of
// passing around a *logger.
var Loggers map[string]*logger

// New returns a pointer to a new logger struct named n.
func New(n string) *logger {
	if Loggers == nil {
		Loggers = make(map[string]*logger)
	}
	newLogger := logger{
		name:     n,
		stdout:   true,
		location: "logs/",
		tlayout:  "2006-01-02 15:04:05 MST",
	}
	Loggers[n] = &newLogger
	return &newLogger
}

// NoStdout disables the logger from going to stdout
func (l *logger) NoStdout() {
	l.stdout = false
	return
}

// Location sets the filepath for the log files
func (l *logger) Location(p string) {
	if p[len(p)-1] != '/' {
		p += "/"
	}
	l.location = p
	return
}

// TimeCode sets the time layout used in the logs
func (l *logger) TimeCode(t string) {
	l.tlayout = t
	return
}

// Wrapper for Error("Info", ...). Shows blue in stdout.
func (l *logger) Info(format string, v ...interface{}) {
	l.Error("Info", Blue, format, v...)
	return
}

// Wrapper for Error("Warning", ...). Shows magenta in stdout.
func (l *logger) Warning(format string, v ...interface{}) {
	l.Error("Warning", Magenta, format, v...)
	return
}

// Wrapper for Error("Fatal", ...). Shows red in stdout.
// After showing and saving the log string, Fatal will
// exit the application with os.Exit(1).
func (l *logger) Fatal(format string, v ...interface{}) {
	l.Error("Fatal", Red, format, v...)
	os.Exit(1)
	return
}

// Error is the core function that will write a log text to file and stdout.
// The log will be of eType type (used for the filename of the log). In
// stdout it will be colored color (see const list). The text will use format
// to Printf v interfaces.
func (l *logger) Error(eType, color, format string, v ...interface{}) {
	if color == "" {
		color = White
	}
	l.writeToStdout(eType, fmt.Sprintf(format, v...), color)
	l.writeToFile(eType, fmt.Sprintf(format, v...))
	return
}

// Write log text to stdout
func (l *logger) writeToStdout(e, s, c string) {
	if !l.stdout {
		return
	}

	now := time.Now().Format(l.tlayout)
	fmt.Printf("%s%s: %s%s: %s%s\n", Grey, now, c, strings.ToUpper(e), Reset, s)
	return
}

// Write log text specific path with filename l.name and path l.location
func (l *logger) writeToFile(e, s string) (n int, err error) {
	if err = checkPath(l.location); err != nil {
		return 0, err
	}

	fileName := l.location + strings.ToLower(e) + ".log"
	errorStr := time.Now().Format(l.tlayout) + ": " + s + "\n"

	saveFile, err1 := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0660)
	if err1 != nil {
		fmt.Printf("%s", err1)
	}

	n, err = saveFile.WriteString(errorStr)
	if err != nil {
		fmt.Printf("%s", err)
	}
	saveFile.Close()
	return
}

// Checks file path to make sure it's available and if not creates it
func checkPath(p string) error {
	_, err := os.Stat(p)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		if err = os.Mkdir(p, 0775); err != nil {
			return fmt.Errorf("%s", "ERROR: Logger - Couldn't create logs folder")
		}
		return nil
	}
	return fmt.Errorf("%s", "ERROR: Logger - Unknown file error")
}
