package http

import (
	"test/internal/dilivery/http/internal"
)

func (s *httpServer) setApi() {
	s.router.POST("/add-money", internal.AddMoney)
}
