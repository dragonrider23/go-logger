package logger

import (
    "fmt"
    "os"
)

func init() {
    if _, err := os.Stat("./logs"); err != nil {
        if os.IsNotExist(err) {
            if err = os.Mkdir("./logs", 0775); err != nil {
                fmt.Print("ERROR: Couldn't create logs folder")
            }
        }
    }
}

type Logger struct {
    name string
}

func New(n string) *Logger {
    return &Logger {
        name: n,
    }
}

func (l *Logger) Info(format string, v ...interface{}) {
    l.writeOut("info", fmt.Sprintf(format, v...))
    return
}

func (l *Logger) Warning(format string, v ...interface{}) {
    l.writeOut("warning", fmt.Sprintf(format, v...))
    return
}

func (l *Logger) Error(format string, v ...interface{}) {
    l.writeOut("error", fmt.Sprintf(format, v...))
    return
}

func (l *Logger) writeOut(e, s string) (n int, err error) {
    fileName := "logs/"+e+".log"
    saveFile, err1 := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0660)
    if err1 != nil {
        fmt.Printf("%s", err1)
    }
    n, err = saveFile.WriteString(s+"\n")
    if err != nil {
        fmt.Printf("%s", err)
    }
    saveFile.Close()
    return
}