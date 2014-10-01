// Copyright 2014 Lee Keitel. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

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
	if !l.file {
		return 0, fmt.Errorf("%s", "Write to file is disabled for this logger")
	}
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
