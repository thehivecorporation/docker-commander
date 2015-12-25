package logger

import (
	"os"

	"github.com/sayden/docker-commander/tests/integration/discovery/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/sayden/docker-commander/tests/integration/discovery/Godeps/_workspace/src/github.com/x-cray/logrus-prefixed-formatter"
)

//Log is a pre-configured logrus logger
var log *logrus.Entry
var logger *logrus.Logger

func init() {
	logger = &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(prefixed.TextFormatter),
		Level:     logrus.DebugLevel,
	}

	log = logger.WithField("prefix", "DOCKER-COMMANDER")
}

func WithField(f interface{}) (_log *logrus.Entry) {
	var prefix string
	if str, ok := f.(string); ok {
		prefix = "DOCKER-COMMANDER:" + str + ":"
	} else {
		prefix = "DOCKER-COMMANDER"
	}

	_log = logger.WithField("prefix", prefix)
	return
}

func Default() *logrus.Entry {
	return log
}
