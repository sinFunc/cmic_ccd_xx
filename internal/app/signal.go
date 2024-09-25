package app

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Signal struct {
	onExits  []func()
	stopFlag chan bool
	osSignal chan os.Signal
	asynFlag bool
}

func (s *Signal) RegisterOnExit(f ...func()) {
	s.onExits = append(s.onExits, f...)
}

func (s *Signal) initSignal() *Signal {
	s.osSignal = make(chan os.Signal, 8)
	signal.Notify(s.osSignal,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGSEGV,
		syscall.SIGABRT)
	return s
}
func (s *Signal) onSignal() {
	for {
		sg := <-s.osSignal
		logger.Infof("received signal: %v", sg.String())
		for _, f := range s.onExits { //like queue fifo
			f()
			// sleep ensures clean-up goroutines being scheduled before os.Exit is called
			time.Sleep(time.Second)
		}
		if s.asynFlag {
			s.stopFlag <- true
		}
		return
	}

}

func (s *Signal) ListenSignal() {
	s.initSignal().onSignal()
}
func (s *Signal) NonBlockListenSignal() {
	s.asynFlag = true
	s.stopFlag <- false
	s.initSignal()

	go s.onSignal()

}
