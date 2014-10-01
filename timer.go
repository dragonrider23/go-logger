// Copyright 2014 Lee Keitel. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package logger

import (
	"strings"
	"time"
)

// timer holds the start time for a log timer.
// Each logger can only have one timer
type timer struct {
	start   time.Time
	running bool
}

// StartTimer associates a timer with logger l
func (l *logger) StartTimer() {
	l.t = timer{
		start:   time.Now(),
		running: true,
	}
	return
}

// StopTimer determines the time elapsed since logger
// l's timer started and issues an Info level log.
// The string "{time}" will be replaced with the
// elapsed time.
func (l *logger) StopTimer(s string) {
	if !l.t.running {
		return
	}

	dateTag := "{time}"
	elapsed := time.Since(l.t.start).String()
	s = strings.Replace(s, dateTag, elapsed, -1)
	l.Info("%s", s)
	l.t.running = false
	return
}
