package server

import (
	"github.com/gin-gonic/gin"
	"mdu/explorer/rest-server/config"
	"strings"
)

var DefaultServer *Server = nil

func init() {
	DefaultServer = &Server{}
	DefaultServer.engine = gin.New()
}

type Server struct {
	engine *gin.Engine
}

func (s *Server) RegisterRoute(method, path string, handler gin.HandlerFunc) {
	switch strings.ToUpper(method) {
	case "GET":
		s.engine.GET(path, handler)
	}
}

func (s *Server) Start() {
	err := s.engine.Run(config.DefaultConfig.HTTPAddress)
	if nil != err {
		panic(err)
	}
}
