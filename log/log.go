package log

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	Debug = &Logger{log.New(ioutil.Discard, "DEBUG ", log.LstdFlags|log.Lshortfile)}
	Info  = &Logger{log.New(os.Stdout, "INFO ", log.LstdFlags|log.Lshortfile)}
	Error = &Logger{log.New(os.Stdout, "ERROR ", log.LstdFlags|log.Lshortfile)}
	Warning = &Logger{log.New(os.Stdout, "WARNING ", log.LstdFlags|log.Lshortfile)}
)

// Logger is a wrapper for log.Logger and provides
// methods to enable and disable logging.
type Logger struct {
	*log.Logger
}

// Disable sets the logging output to /dev/null.
func (l *Logger) Disable() {
	l.SetOutput(ioutil.Discard)
}

// Enable sets the logging output to stdout.
func (l *Logger) Enable() {
	l.SetOutput(os.Stdout)
}
