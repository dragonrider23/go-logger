package logger

import (
    "testing"
 //   "fmt"
)

func TestLogger(t *testing.T) {
    var log = New("api")

    log.Info("%s", "Hello")
}