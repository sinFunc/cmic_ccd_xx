package server

import (
	"cmic_ccd_xx/cmd/config"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

func init() {
	config.RegisterLoggerInitializer(func(l *logrus.Logger) {
		logger = l.WithField("module", "service")
	})
}

type InitData struct {
	LocalIp   string
	LocalPort int
}

type Base interface {
	Init(data *InitData) error
	Start() error
	Stop() error
}
