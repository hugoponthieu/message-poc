package http

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	port        string
	allowOrigin string
	Engine      *gin.Engine
}

func NewServer(port string, allowOrigin string) *Server {
	return &Server{
		port:        port,
		Engine:      gin.Default(),
		allowOrigin: allowOrigin,
	}
}

func (s *Server) Start() error {
	return s.Engine.Run(":" + s.port)
}
