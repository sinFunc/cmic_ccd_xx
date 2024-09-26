package server

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

type WebsocketServer struct {
	HttpServer
	conn     *websocket.Conn
	upgrater *websocket.Upgrader
}

func (s *WebsocketServer) Init() error {
	s.upgrater = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	return nil
}

func (s *WebsocketServer) Start() (err error) {
	if s.localPort <= 0 {
		e := fmt.Sprintf("port is invaild.%v", s.localPort)
		return errors.New(e)
	}

	url := s.localIp + ":" + strconv.FormatInt(int64(s.localPort), 10)

	s.server = &http.Server{
		Addr:    url,
		Handler: nil,
	}

	//register many functions
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		s.upGrader(writer, request)
	})
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		logger.Infof("test-host=%v", request.Host)
	})

	logger.Infof("Start websocket server listen with url=%v%v", url, s.pattern)

	err = s.server.ListenAndServe() //block
	return nil
}

func (s *WebsocketServer) upGrader(writer http.ResponseWriter, request *http.Request) {
	conn, err := s.upgrater.Upgrade(writer, request, nil)
	if err != nil {
		logger.Errorf("can not upgrate http to websocket")
		conn.Close()
		return
	}

	logger.Infof("create websocket one connection")

	conn.SetCloseHandler(func(code int, text string) error {
		logger.Infof("closehandler of websocket conn with code=%v,text=%v", code, text)
		return nil
	})
	conn.SetPongHandler(func(appData string) error {
		logger.Infof("closehandler of websocket conn with appData=%v", appData)
		return nil
	})
	conn.SetPingHandler(func(appData string) error {
		logger.Infof("closehandler of websocket conn with appData=%v", appData)
		return nil
	})

	s.conn = conn

}

func (s *WebsocketServer) Stop() error {
	if s.conn != nil {
		logger.Infof("stop websocket server listening")
		s.conn.Close()
	}

	return s.HttpServer.Stop()
}

func (s *WebsocketServer) CandidateHandler(writer http.ResponseWriter, request *http.Request) {
	
}

func (s *WebsocketServer) rcvHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("websocketServer")
	//logger.Info(request.Body)
}

func (s *WebsocketServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("websocketServer")
}
