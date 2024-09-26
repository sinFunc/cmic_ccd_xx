package server

import (
	"cmic_ccd_xx/cmd/config"
	"errors"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

func init() {
	config.RegisterLoggerInitializer(func(l *logrus.Logger) {
		logger = l.WithField("module", "service")
	})
}

// go does not support complex inhert
type Base interface {
	Init() error
	Start() error
	Stop() error
	Destroy()
}

// define common fields and methods
type BaseServer struct {
	localIp   string
	localPort int
	pattern   string            //for http homepage url
	params    map[string]string //extra params for expending
}

func (b *BaseServer) Init() error {
	//just use default value
	return nil
}

func (b *BaseServer) AddOneExtraParam(k string, v string) *BaseServer {
	if b.params == nil {
		b.params = make(map[string]string)
	}
	//rewrite policy
	b.params[k] = v

	return b
}

func (b *BaseServer) SetLocalIp(lip string) *BaseServer {
	b.localIp = lip
	return b
}
func (b *BaseServer) SetLocalPort(port int) *BaseServer {
	b.localPort = port
	return b
}
func (b *BaseServer) SetPattern(pattern string) *BaseServer {
	b.pattern = pattern
	return b
}

func (b *BaseServer) Start() error {
	return errors.New("Please override this ")
}

func (b *BaseServer) Stop() error {
	return errors.New("Please override this ")
}

func (b *BaseServer) Destroy() {
	b.params = nil
}
