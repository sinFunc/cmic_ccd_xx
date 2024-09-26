package server

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type HttpServer struct {
	BaseServer
	server *http.Server
}

//func (s *HttpServer) Init() error {
//	fmt.Printf("httpServer...")
//	return nil
//}

func (s *HttpServer) Start() (err error) {
	if s.localPort <= 0 {
		e := fmt.Sprintf("port is invaild.%v", s.localPort)
		return errors.New(e)
	}

	url := s.localIp + ":" + strconv.FormatInt(int64(s.localPort), 10)

	s.server = &http.Server{
		Addr:    url,
		Handler: &HttpServer{},
	}

	//pattern := "/"
	//if s.pattern != "" {
	//	pattern = s.pattern
	//}

	//parse params if necessary

	//http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
	//	s.rcvHandler(writer, request)
	//})

	logger.Infof("Start http server listen with url=%v%v", url, s.pattern)

	err = s.server.ListenAndServe() //block

	return

}

func (s *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("ServeHTTP------httpServer%V", request.URL.String())

}

func (s *HttpServer) Stop() (err error) {
	if s.server != nil {
		logger.Infof(":Close http server connection")
		s.server.Close()
	}

	return

}

func (s *HttpServer) rcvHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("httpServer")
	//logger.Info(request.Body)
}
