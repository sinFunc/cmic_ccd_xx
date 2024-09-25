package config

import (
	"errors"
	"github.com/sinFunc/singleton"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"strconv"
)

// global logger handler
var pLogger *logrus.Logger

type loggerInitFunc func(l *logrus.Logger)

var loggerInitializers []loggerInitFunc

func RegisterLoggerInitializer(init loggerInitFunc) {
	loggerInitializers = append(loggerInitializers, init)
}

func initLogger() error {
	if pLogger != nil {
		return errors.New("already")
	}

	pLogger = logrus.New()
	pLogger.SetOutput(os.Stdout)

	l := singleton.GetInstance[AppConfig]().(*AppConfig).Logger
	fc := true //file does not support color
	if l.File != "" {
		fc = false
		file, err := os.OpenFile(l.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		pLogger.SetOutput(file)
	}
	if l.Level != "" {
		pLogger.SetLevel(getLoglevel(l.Level))
	}

	pLogger.SetReportCaller(true)
	pLogger.SetFormatter(&logrus.TextFormatter{
		ForceColors: fc,
		//TimestampFormat: "2006-01-02 15:02:03", //bug ts does not change
		FullTimestamp: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			f := path.Base(frame.File) + "-" + strconv.FormatInt(int64(frame.Line), 10)
			return frame.Function, f
		},
	})

	for _, i := range loggerInitializers {
		i(pLogger)
	}

	return nil
}

func getLoglevel(level string) (logLevel logrus.Level) {
	switch level {
	case "trace":
		logLevel = logrus.TraceLevel
	case "debug":
		logLevel = logrus.DebugLevel
	case "warn":
		logLevel = logrus.WarnLevel
	case "error":
		logLevel = logrus.ErrorLevel
	case "fatal":
		logLevel = logrus.FatalLevel
	case "panic":
		logLevel = logrus.PanicLevel
	case "info":
		fallthrough
	default:
		logLevel = logrus.InfoLevel
	}
	return
}
