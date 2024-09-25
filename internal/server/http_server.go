package server

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type HttpServer struct {
	server *http.Server
	ip     string
	port   int
}

func (s *HttpServer) Init(d *InitData) (err error) {
	s.ip = d.LocalIp
	s.port = d.LocalPort

	return
}

func (s *HttpServer) Start() (err error) {
	if s.ip == "" || s.port <= 0 {
		e := fmt.Sprintf("ip:port is invaild.%v:%v", s.ip, s.port)
		return errors.New(e)
	}
	url := s.ip + ":" + strconv.FormatInt(int64(s.port), 10)

	s.server = &http.Server{
		Addr:    url,
		Handler: nil,
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		s.rcvHandler(writer, request)
	})

	logger.Infof("Start http server listen")

	err = s.server.ListenAndServe() //block

	return

}

func (s *HttpServer) Stop() (err error) {
	if s.server != nil {
		s.server.Close()
	}

	return

}

func (s HttpServer) RegisterHandler(name string, handler func()) {

}

func (s *HttpServer) rcvHandler(writer http.ResponseWriter, request *http.Request) {
	logger.Info(request.Body)
}
