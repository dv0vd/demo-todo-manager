package logger

import "github.com/sirupsen/logrus"

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.SetLevel(logrus.DebugLevel)
	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}
