package web

import "github.com/gin-gonic/gin"

type HttpServer struct {
	route *gin.Engine
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		route: NewRouter(),
	}
}
func (s *HttpServer) Run() error {

	if err := s.route.Run(":" + "8080"); err != nil {
		return err
	}
	return nil
}
