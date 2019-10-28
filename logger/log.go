package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// init ...
func init() {
	log.SetFormatter(&prefixed.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})

	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

// Info
func Info(args ...interface{}) {
	log.Info(args)
}

// Debug
func Debug(args ...interface{}) {
	log.Debug(args)
}

// Fatal
func Fatal(args ...interface{}) {
	log.Fatal(args)
}

// Info
func Error(args ...interface{}) {
	log.Error(args)
}
