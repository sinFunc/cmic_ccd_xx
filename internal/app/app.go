package app

import (
	"cmic_ccd_xx/cmd/config"
	"cmic_ccd_xx/internal/server"
	"github.com/sinFunc/singleton"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

func init() {
	config.RegisterLoggerInitializer(func(l *logrus.Logger) {
		logger = l.WithField("module", "app")
	})
}

type App struct {
	sig    *Signal
	cancel func()
	server server.Base
}

func (a *App) initConfig() {
	const filePath = "../configs/config.yml"

	ac := singleton.GetInstance[config.AppConfig]().(*config.AppConfig)
	if err := ac.Init(filePath); err != nil {
		logger.Fatalf("Initialize config failed:%v\n", err)
		return
	}

	logger.Infof("Initialize config from file:%v\n", filePath)

}
func (a *App) listenSignal() {
	a.sig = &Signal{}

	s1 := func() {
		logger.Infof("s1")
	}
	s2 := func() {
		logger.Infof("s2")
	}

	a.sig.RegisterOnExit(s1, s2)
	a.sig.ListenSignal()
	//s.NonBlockListenSignal()

}

func (a *App) start() {
	a.startService()
}

// start network service
func (a *App) startService() {
	s := &server.HttpServer{}
	a.server = s
	d := &server.InitData{
		LocalIp:   "127.0.0.1",
		LocalPort: 7777,
	}
	if s.Init(d) != nil {
		return
	}
	go s.Start()

}

func (a *App) stop() {
	if a.server != nil {
		a.server.Stop()
	}
}

// interface to start app
func Start() {
	a := &App{}
	a.initConfig()
	a.start()
	a.listenSignal()

	a.stop()

}
