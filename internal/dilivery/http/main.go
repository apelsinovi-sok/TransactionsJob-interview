package http

import (
	"github.com/gin-gonic/gin"
	"test/internal/core/interfaces"
)

type httpServer struct {
	router     *gin.Engine
	repository interfaces.TransferRepository
}

func New(repository interfaces.TransferRepository) *httpServer {
	server := &httpServer{router: gin.Default(), repository: repository}
	server.setExternalParamInHandlers()
	server.setApi()
	return server
}

func (s *httpServer) setExternalParamInHandlers() {
	s.router.Use(func(c *gin.Context) {
		c.Set("repository", &s.repository)
		c.Next()
	})
}

func (s *httpServer) Run() {
	err := s.router.Run()
	if err != nil {
		panic(err.Error())
	}
}
