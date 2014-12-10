// Copyright 2014 Lee Keitel. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"os"
)

// Colors that can be used for Error()
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Grey    = "\x1B[90m"
)

var verbosity int = 2

// Type logger is the struct returned and used for logging
// The user can set its properties using the associated functions.
type logger struct {
	name, location, tlayout string
	stdout, file, raw       bool
	verbosity               int
	t                       timer
}

// loggers is the collection of all logger types used.
// A logger can be called by their names instead of
// passing around a *logger.
var loggers map[string]*logger

// New returns a pointer to a new logger struct named n.
func New(n string) *logger {
	// Initialize loggers
	if loggers == nil {
		loggers = make(map[string]*logger)
	}
	// If logger with n name is already created, return it
	if loggers[n] != nil {
		return loggers[n]
	}
	// Create new logger
	newLogger := &logger{
		name:      n,
		stdout:    true,
		file:      true,
		raw:       false,
		verbosity: -1,
		location:  "logs/",
		tlayout:   "2006-01-02 15:04:05 MST",
	}
	// Add to loggers
	loggers[n] = newLogger
	return newLogger
}

// Get retrives the logger with name n
func Get(n string) *logger {
	log := loggers[n]
	// If logger doesn't exist, create it
	if loggers[n] == nil {
		log = New(n)
	}
	return log
}

// Set global verbosity for stdout
func Verbose(v int) {
	if v < 0 {
		v = 0
	} else if v > 3 {
		v = 3
	}
	verbosity = v
	return
}

// Gets current global verbosity
func GetVerboseLevel() int {
	return verbosity
}

// Set verbose level of indivindual logger
// -1 means use the global verbosity
func (l *logger) Verbose(v int) *logger {
	if v < -1 {
		v = -1
	} else if v > 3 {
		v = 3
	}
	l.verbosity = v
	return l
}

// NoStdout disables the logger from going to stdout
func (l *logger) NoStdout() *logger {
	l.stdout = false
	return l
}

// NoFile disables the logger from writting to a log file
func (l *logger) NoFile() *logger {
	l.file = false
	return l
}

// Raw tells the logger writer to not pre-include the date
func (l *logger) Raw() *logger {
	l.raw = true
	return l
}

// Path sets the filepath for the log files
func (l *logger) Path(p string) *logger {
	if p[len(p)-1] != '/' {
		p += "/"
	}
	l.location = p
	return l
}

// TimeCode sets the time layout used in the logs
func (l *logger) TimeCode(t string) *logger {
	l.tlayout = t
	return l
}

// Remove logger l from the loggers map
func (l *logger) Close() {
	delete(loggers, l.name)
	return
}

// Wrapper for Log("Info", ...). Shows blue in stdout.
func (l *logger) Info(format string, v ...interface{}) {
	l.Log("Info", Cyan, format, v...)
	return
}

// Wrapper for Log("Warning", ...). Shows magenta in stdout.
func (l *logger) Warning(format string, v ...interface{}) {
	l.Log("Warning", Magenta, format, v...)
	return
}

// Wrapper for Log("Error", ...). Shows red in stdout.
func (l *logger) Error(format string, v ...interface{}) {
	l.Log("Error", Red, format, v...)
	return
}

// Wrapper for Log("Fatal", ...). Shows red in stdout.
// After showing and saving the log string, Fatal will
// exit the application with os.Exit(1).
func (l *logger) Fatal(format string, v ...interface{}) {
	l.Log("Fatal", Red, format, v...)
	os.Exit(1)
	return
}

// Log is the core function that will write a log text to file and stdout.
// The log will be of eType type (used for the filename of the log). In
// stdout it will be colored color (see const list). The text will use format
// to Printf v interfaces.
func (l *logger) Log(eType, color, format string, v ...interface{}) {
	if color == "" {
		color = White
	}
	l.writeAll(eType, fmt.Sprintf(format, v...), color)
	return
}
