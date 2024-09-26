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
		a.stop()
	}

	a.sig.RegisterOnExit(s1)
	a.sig.ListenSignal()
	//s.NonBlockListenSignal()

}

// start service
func (a *App) start() {
	a.startService()
	a.listenSignal()
}

// start network service
func (a *App) startService() {
	//s := &server.HttpServer{}
	s := &server.WebsocketServer{}
	a.server = s
	if s.Init() != nil {
		logger.Errorf("initialize websocket failed")
		return
	}

	c := singleton.GetInstance[config.AppConfig]().(*config.AppConfig).Server
	s.SetLocalIp(c.Ip).SetLocalPort(c.Port).SetPattern("/")

	go s.Start()

}

func (a *App) stop() {
	if a.server != nil {
		logger.Infof("try to stop app...")
		a.server.Stop()
	}
}

// interface to start app
func Start() {
	a := &App{}
	a.initConfig()
	a.start()
	//a.listenSignal()
	//a.stop()

}
